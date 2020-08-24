package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// wsHandler handles requests to do with opening or managing Websocket
// connections (any /ws/* endpoints).
type wsHandler struct{}

// newWSHandler returns a pointer to a new WSHandler.
func newWSHandler() *wsHandler {
	return &wsHandler{}
}

// defaultHandler serves requests at the /ws route. Creates a new game (or adds
// a new player to the existing game that corresponds with the provided gameID),
// and attempts to establish a Websocket connection with a client. Expects
// "gameID" as a query parameter when the /ws endpoint is hit.
func (h *wsHandler) defaultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: temporary! just for testing
	fmt.Fprintln(w, "hello from the ws handler :)")
}

// registerRoutes registers handlers for all of the routes that wsHandler
// supports.
func (h *wsHandler) registerRoutes(router *mux.Router) {
	router.HandleFunc("/ws", h.defaultHandler)
}
