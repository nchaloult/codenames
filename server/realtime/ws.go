package realtime

import "github.com/gorilla/websocket"

type eventKind string

const (
	changeDisplayName eventKind = "changeDisplayName"
)

// event mirrors the structure of JSON messages that clients send to the server
// via a Websocket connection.
type event struct {
	kind eventKind
	body interface{}
}

// eventResponse mirrors the structure of JSON messages that the server sends in
// response to a client's Websocket message.
type eventResponse struct {
	ok   bool
	kind eventKind
	body interface{}
}

// constructAndSendResponse builds an eventResponse struct with the provided
// fields, marshals it to JSON, and sends it along the provided Websocket
// connection.
func constructAndSendResponse(
	conn *websocket.Conn,
	kind eventKind,
	body interface{},
) {
	response := eventResponse{ok: true, kind: kind, body: body}
	conn.WriteJSON(response)
}

// constructAndSendErr builds an eventResponse struct with the provided error,
// marshals it to JSON, and sends it along the provided Websocket connection.
func constructAndSendErr(conn *websocket.Conn, err error) {
	response := eventResponse{ok: false, body: err}
	conn.WriteJSON(response)
}
