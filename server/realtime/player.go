package realtime

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Player stores information about a player in a Game, as well as the Websocket
// connection that they're connected to the server with. Players are managed by
// an Interactor.
type Player struct {
	Conn        *websocket.Conn
	ID          string
	DisplayName string
	IsOnRedTeam bool
	IsSpymaster bool
}

// NewPlayer returns a pointer to a new Player object with the provided
// connection pointer. Generates a new unique Player ID and sends it to the
// client listening on the other end of the Websocket connection. Sets
// IsOnRedTeam to true and IsSpymaster to false by default.
func NewPlayer(conn *websocket.Conn) *Player {
	id := uuid.New()
	playerIDMsg := map[string]interface{}{
		"id": id,
	}
	conn.WriteJSON(playerIDMsg)

	return &Player{
		Conn:        conn,
		ID:          id.String(),
		IsOnRedTeam: true,
		IsSpymaster: false,
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
			return
		}

		// Decide how to behave depending on the event's type/kind.
		switch event.Kind {
		case changeDisplayName:
			p.DisplayName = event.Body.(string)
		default:
			errMsg := fmt.Errorf("unrecognized eventKind: %v", event.Kind)
			constructAndSendErr(p.Conn, errMsg)
		}
	}
}
