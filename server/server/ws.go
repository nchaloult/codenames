package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// WSHandler handles requests to do with opening or managing Websocket
// connections (any /ws/* endpoints).
type WSHandler struct{}

// NewWSHandler returns a pointer to a new WSHandler.
func NewWSHandler() *WSHandler {
	return &WSHandler{}
}

// defaultHandler serves requests at the /ws route. Creates a new game (or adds
// a new player to the existing game that corresponds with the provided gameID),
// and attempts to establish a Websocket connection with a client. Expects
// "gameID" as a query parameter when the /ws endpoint is hit.
func (h *WSHandler) defaultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: temporary! just for testing
	fmt.Fprintln(w, "hello from the ws handler :)")
}

// RegisterRoutes registers handlers for all of the routes that wsHandler
// supports.
func (h *WSHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ws", h.defaultHandler)
}
