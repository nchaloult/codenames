package realtime

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Artbitrarily chosen value. Regardless of how many Players are (realistically)
// in a Game, I think it's pretty unlikely that 16 of those Players will all
// need to broadcast an event at nearly the same time.
const broadcastedMsgsBufferSize = 16

// Player stores information about a player in a Game, as well as the Websocket
// connection that they're connected to the server with. Players are managed by
// an Interactor.
type Player struct {
	Conn        *websocket.Conn
	ID          string
	DisplayName string
	IsOnRedTeam bool
	IsSpymaster bool
	// Buffered channel of events that were broadcasted from other Players.
	broadcastedMsgs chan *event
}

// NewPlayer returns a pointer to a new Player object with the provided
// connection pointer. Sets IsOnRedTeam to true and IsSpymaster to false by
// default.
//
// NewPlayer sends a NewPlayerID event to the client once a UUID has been
// generated for them.
func NewPlayer(conn *websocket.Conn) *Player {
	id := uuid.New().String()
	eventBody := map[string]string{"id": id}
	ConstructAndSendEvent(conn, NewPlayerID, eventBody)
	return &Player{
		Conn:            conn,
		broadcastedMsgs: make(chan *event, broadcastedMsgsBufferSize),
		ID:              id,
		IsOnRedTeam:     true,
		IsSpymaster:     false,
	}
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
			// TODO: remove this Player from their Interactor.
			return
		}

		// Decide how to behave depending on the event's type/kind.
		switch event.Kind {
		case ChangeDisplayName:
			p.DisplayName = event.Body.(string)
		default:
			errMsg := fmt.Errorf("unrecognized eventKind: %v", event.Kind)
			constructAndSendErr(p.Conn, errMsg)
		}
	}
}
