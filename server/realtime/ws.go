package realtime

import "github.com/gorilla/websocket"

// EventKind represents the type of an event or eventResponse that's sent over a
// Websocket connection.
type EventKind string

const (
	// LobbyInfo describes an event that the server sends when a client first
	// connects to a /:gameID URL. It communicates whether a game with the
	// provided ID already exists or not.
	LobbyInfo EventKind = "LOBBY_INFO"
	// ChangeDisplayName describes an event that the client sends when they
	// change their display name.
	ChangeDisplayName EventKind = "CHANGE_DISPLAY_NAME"
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
