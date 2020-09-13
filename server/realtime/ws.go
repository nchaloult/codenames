package realtime

import "github.com/gorilla/websocket"

// EventKind represents the type of an event or eventResponse that's sent over a
// Websocket connection.
type EventKind string

const (
	// LobbyInfo describes an event that the server sends when a client first
	// connects to a /:gameID URL. It communicates whether a game with the
	// provided ID already exists or not, as well as the other players that are
	// already in that game.
	LobbyInfo EventKind = "LOBBY_INFO"
	// NewPlayerID describes an event that the server sends to a client after
	// making a new Player object for them. Clients need to know what their
	// Player IDs are in order to include in the bodies of select events.
	NewPlayerID EventKind = "NEW_PLAYER_ID"
	// NewPlayerJoined describes an event that a client sends after it first
	// connects to a /:gameID URL and sets its display name. This gets
	// broadcasted to all other Players in that game lobby so that they may add
	// this new player to their red teams client-side.
	NewPlayerJoined EventKind = "NEW_PLAYER_JOINED"
	// ChangeDisplayName describes an event that the client sends when they
	// change their display name.
	ChangeDisplayName EventKind = "CHANGE_DISPLAY_NAME"
	// SomeoneElseChangeDisplayName describes an event that the server sends to
	// all clients in a game if one of the players changes their display name.
	SomeoneElseChangeDisplayName EventKind = "SOMEONE_ELSE_CHANGE_DISPLAY_NAME"
	// ChangeTeam describes an event that the client has swapped teams in a game
	// lobby.
	ChangeTeam EventKind = "CHANGE_TEAM"
	// SomeoneElseChangeTeam describes an event that some other Player in a game
	// lobby changed teams.
	SomeoneElseChangeTeam EventKind = "SOMEONE_ELSE_CHANGE_TEAM"
)

// event mirrors the structure of JSON messages that clients send to the server
// via a Websocket connection.
type event struct {
	Kind EventKind   `json:"kind"`
	Body interface{} `json:"body"`
}

// eventResponse mirrors the structure of JSON messages that the server sends in
// response to a client's Websocket message.
type eventResponse struct {
	Ok   bool        `json:"ok"`
	Kind EventKind   `json:"kind"`
	Body interface{} `json:"body"`
}

// broadcastMsg objects are stored in an Interactor's broadcast channel. They
// contain an event to be broadcasted to all Players in a Game except for the
// Player who originally sent the event.
type broadcastMsg struct {
	// ID of the Player who originally sent the event. This is so that Players
	// who send events that are broadcasted to everyone in a Game don't receive
	// their own event.
	originClientID string
	// The event to broadcast to all other Players in a Game.
	event *event
}

// ConstructAndSendEvent builds an event struct with the provided fields,
// marshals it to JSON, and sends it along the provided Websocket connection.
func ConstructAndSendEvent(conn *websocket.Conn, kind EventKind, body interface{}) {
	event := event{Kind: kind, Body: body}
	conn.WriteJSON(event)
}

// constructAndSendResponse builds an eventResponse struct with the provided
// fields, marshals it to JSON, and sends it along the provided Websocket
// connection. Body parameter may be nil.
func constructAndSendResponse(
	conn *websocket.Conn,
	kind EventKind,
	body interface{},
) {
	response := eventResponse{Ok: true, Kind: kind}
	if body != nil {
		response.Body = body
	}
	conn.WriteJSON(response)
}

// constructAndSendErr builds an eventResponse struct with the provided error,
// marshals it to JSON, and sends it along the provided Websocket connection.
func constructAndSendErr(conn *websocket.Conn, err error) {
	response := eventResponse{Ok: false, Body: err}
	conn.WriteJSON(response)
}
