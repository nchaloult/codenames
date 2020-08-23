package model

import (
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
