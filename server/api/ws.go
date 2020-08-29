package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/nchaloult/codenames/model"
	"github.com/nchaloult/codenames/realtime"
)

const (
	devClientURL  = "localhost:3000"
	prodClientURL = "foobar.com" // TODO: change this once you deploy front-end.
)

// WSHandler handles requests to do with opening or managing Websocket
// connections (any /ws/* endpoints).
type WSHandler struct {
	manager *realtime.Manager
}

// NewWSHandler returns a pointer to a new WSHandler initialized with the
// provided fields.
func NewWSHandler(manager *realtime.Manager) *WSHandler {
	return &WSHandler{manager}
}

// checkClientOrigin makes sure that the client application that's trying to
// establish a websocket connection with us is someone we recognize: either the
// client's origin in development (localhost:some-port) or the client's
// production URL.
func checkClientOrigin(r *http.Request) bool {
	origin := r.Header["Origin"]
	if len(origin) == 0 {
		return true
	}

	u, err := url.Parse(origin[0])
	if err != nil {
		return false
	}

	// TODO: move these two whitelisted URLs to env vars, maybe?
	return u.Host == devClientURL || u.Host == prodClientURL
}

// defaultHandler serves requests at the /ws route. Creates a new game (or adds
// a new player to the existing game that corresponds with the provided gameID),
// and attempts to establish a Websocket connection with a client. Expects
// "gameID" as a query parameter when the /ws endpoint is hit.
func (h *WSHandler) defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Extract info from query params.
	queryParams := r.URL.Query()
	if _, ok := queryParams["gameID"]; !ok {
		errMsg := "the \"gameID\" query param is required"
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	gameID := queryParams["gameID"][0]

	// Look for an active game associated with the provided gameID. If one
	// doesn't exist, then create a new one.
	if _, ok := h.manager.ActiveGames[gameID]; !ok {
		// TODO: Write function somewhere in model package to randomly generate
		// a dictionary. For now, I'm just using a dummy/stubbed one.
		dictionary := [25]string{"foo", "bar", "baz"}
		newGame := model.NewGame(dictionary)
		newInteractor := realtime.NewInteractor(newGame)
		h.manager.ActiveGames[gameID] = newInteractor
	}

	upgrader := websocket.Upgrader{}
	// Set up upgrader to only accept connections from a few specified origins.
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return checkClientOrigin(r)
	}

	// Attempt to establish a websocket connection with the client.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errMsg := fmt.Sprintf("failed to upgrade to a websocket connection: %v",
			err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// Create a new player for our connected client.
	// TODO: make this async by pushing this newPlayer pointer to some channel.
	// This will break if more than one client hits the /ws endpoint at once
	// with the same gameID.
	newPlayer := realtime.NewPlayer(conn)
	h.manager.ActiveGames[gameID].Players[newPlayer.DisplayName] = newPlayer

	go newPlayer.ListenForEvents()
}

// RegisterRoutes registers handlers for all of the routes that wsHandler
// supports.
func (h *WSHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ws", h.defaultHandler)
}
