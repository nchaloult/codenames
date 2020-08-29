package realtime

import "github.com/gorilla/websocket"

type eventKind string

const (
	changeDisplayName eventKind = "changeDisplayName"
)

// event mirrors the structure of JSON messages that clients send to the server
// via a Websocket connection.
type event struct {
	Kind eventKind   `json:"kind"`
	Body interface{} `json:"body"`
}

// eventResponse mirrors the structure of JSON messages that the server sends in
// response to a client's Websocket message.
type eventResponse struct {
	Ok   bool        `json:"ok"`
	Kind eventKind   `json:"kind"`
	Body interface{} `json:"body"`
}

// constructAndSendResponse builds an eventResponse struct with the provided
// fields, marshals it to JSON, and sends it along the provided Websocket
// connection.
func constructAndSendResponse(
	conn *websocket.Conn,
	kind eventKind,
	body interface{},
) {
	response := eventResponse{Ok: true, Kind: kind, Body: body}
	conn.WriteJSON(response)
}

// constructAndSendErr builds an eventResponse struct with the provided error,
// marshals it to JSON, and sends it along the provided Websocket connection.
func constructAndSendErr(conn *websocket.Conn, err error) {
	response := eventResponse{Ok: false, Body: err}
	conn.WriteJSON(response)
}
