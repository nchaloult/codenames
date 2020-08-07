package model

import "github.com/gorilla/websocket"

// Player stores information about a player in a Game, as well as the Websocket
// connection that they're connected to the server with. Players are managed by
// an Interactor.
type Player struct {
	Conn        *websocket.Conn
	DisplayName string
	IsOnRedTeam bool
	IsSpymaster bool
}

// NewPlayer returns a pointer to a new Player object with initialized fields.
// Sets IsOnRedTeam to true and IsSpymaster to false by default.
func NewPlayer(conn *websocket.Conn, displayName string) *Player {
	return &Player{conn, displayName, true, false}
}
