package realtime

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Arbitrarily chosen value. Regardless of how many Players are (realistically)
// in a Game, I think it's pretty unlikely that 16 of those Players will all
// need to broadcast an event at nearly the same time.
const broadcastedMsgsBufferSize = 16

// Player stores information about a player in a Game, as well as the Websocket
// connection that they're connected to the server with. Players are managed by
// an Interactor.
type Player struct {
	Conn       *websocket.Conn
	interactor *Interactor
	// Buffered channel of events that were broadcasted from other Players.
	broadcastedMsgs chan *event

	ID          string
	DisplayName string
	IsOnRedTeam bool
	IsSpymaster bool
}

// NewPlayer returns a pointer to a new Player object with the provided
// connection pointer. Sets IsOnRedTeam to true and IsSpymaster to false by
// default.
//
// NewPlayer sends a NewPlayerID event to the client once a UUID has been
// generated for them.
func NewPlayer(conn *websocket.Conn, i *Interactor) *Player {
	id := uuid.New().String()
	eventBody := map[string]string{"id": id}
	ConstructAndSendEvent(conn, NewPlayerID, eventBody)
	return &Player{
		Conn:            conn,
		interactor:      i,
		broadcastedMsgs: make(chan *event, broadcastedMsgsBufferSize),
		ID:              id,
		IsOnRedTeam:     true,
		IsSpymaster:     false,
	}
}

// broadcastToOtherPlayers constructs a broadcastMsg struct and pushes it onto
// an Interactor's msgsToBroadcast channel.
func (p *Player) broadcastToOtherPlayers(kind EventKind, body interface{}) {
	event := event{Kind: kind, Body: body}
	msg := &broadcastMsg{originClientID: p.ID, event: &event}
	p.interactor.msgsToBroadcast <- msg
}

// SendBroadcastedEventsToClient watches a Player's broadcastedMsgs channel for
// messages from an Interactor, and sends it to our client along our WS
// connection.
func (p *Player) SendBroadcastedEventsToClient() {
	defer func() {
		p.Conn.Close()
	}()
	for broadcastedEvent := range p.broadcastedMsgs {
		ConstructAndSendEvent(p.Conn, broadcastedEvent.Kind, broadcastedEvent.Body)
	}
	// Once we reach this point, the channel's been closed. Send a CloseMessage.
	// This isn't done in the above deferred func because there are situations
	// where we want to return from this func, but not close the WS connection.
	// For instance, if an error is returned when we try to write a message to
	// the WS connection.
	p.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}

// ListenForEvents watches a Player's Websocket connection for messages from
// their client, like if they click a button, for example.
func (p *Player) ListenForEvents() {
	defer func() {
		// TODO: remove this Player from their Interactor's list of Players.
		// Maybe push this Player's ID onto a channel or something.
		p.Conn.Close()
	}()

	log.Printf("Player %s is listening for events...", p.ID)

	for {
		event := event{}
		err := p.Conn.ReadJSON(&event)
		if err != nil {
			log.Printf("Player %s websocket unexpected error: %v", p.ID, err)
			p.Conn.Close()
			// Remove this Player from their Interactor.
			delete(p.interactor.Players, p.ID)
			p.interactor = nil
			return
		}

		// Decide how to behave depending on the event's type/kind.
		switch event.Kind {
		case ChangeDisplayName:
			p.DisplayName = event.Body.(string)
		case ChangeTeam:
			switchingToRedTeam := event.Body.(bool)
			if (switchingToRedTeam && !p.IsOnRedTeam) || (!switchingToRedTeam && p.IsOnRedTeam) {
				p.IsOnRedTeam = switchingToRedTeam
				body := map[string]interface{}{
					"id":          p.ID,
					"displayName": p.DisplayName,
					"isOnRedTeam": switchingToRedTeam,
				}
				p.broadcastToOtherPlayers(SomeoneElseChangeTeam, body)
			}
		default:
			errMsg := fmt.Errorf("unrecognized eventKind: %v", event.Kind)
			constructAndSendErr(p.Conn, errMsg)
		}
	}
}
