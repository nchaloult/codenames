package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/nchaloult/codenames/model"
	"github.com/nchaloult/codenames/realtime"
)

const (
	minPort       = 1025
	maxPort       = 65535
	devClientURL  = "localhost:3000"
	prodClientURL = "foobar.com" // TODO: change this once you deploy front-end.
)

// Server exposes HTTP API endpoints that let clients mutate and interact with
// a Game's state, serves the static files for the web front-end, and maintains
// a collection of active Games in memory.
type Server struct {
	// port is the port number that an HTTP server will listen for requests on.
	port int

	// activeGames stores pointers to Interactors for all ongoing games that the
	// server is handling, indexed by gameIDs.
	activeGames map[string]*realtime.Interactor
}

// NewServer returns a pointer to a new Server object that's configured with the
// provided port number.
func NewServer(port int) (*Server, error) {
	if port < minPort || port > maxPort {
		return nil, fmt.Errorf(
			"the provided port must be within the range: [%d, %d]",
			minPort,
			maxPort,
		)
	}

	return &Server{
		port:        port,
		activeGames: make(map[string]*realtime.Interactor, 0),
	}, nil
}

// routeHandler describes objects which handle requests to specific HTTP
// endpoints.
type routeHandler interface {
	RegisterRoutes(*mux.Router)
}

// Start registers all of the HTTP handler funcs with their corresponding routes
// and begins listening for new requests that come in on those routes.
func (s *Server) Start() {
	// Listen for OS interrupts (like if someone presses Ctrl+C or something).
	// Spin down gracefully if this process is interrupted.
	sigintChan := make(chan os.Signal, 1)
	signal.Notify(sigintChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
		os.Kill,
	)
	go func() {
		<-sigintChan

		log.Println("Spinning down....")
		os.Exit(0)
	}()

	router := mux.NewRouter()

	// Register routes with their corresponding handler funcs.
	router.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	router.HandleFunc("/ws", s.wsHandler)

	// Stand up the server.
	log.Printf("Listening on port %d....\n", s.port)
	portAddr := fmt.Sprintf(":%d", s.port)
	log.Fatal(http.ListenAndServe(portAddr, router))
}

// healthHandler serves requests at the /health route. Responds with information
// about the server's state.
//
// TODO: include health information about the database in response.
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"ok":             true,
		"numActiveGames": len(s.activeGames),
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		errCode := http.StatusInternalServerError
		errMsg := fmt.Sprintf("Failed to encode response as JSON: %v\n", err)
		http.Error(w, errMsg, errCode)
		return
	}
}

// wsRequestBody represents the structure of JSON bodies in POST requests to
// the /ws route.
type wsRequestBody struct {
	gameID      string
	displayName string
}

// wsHandler serves requests at the /ws route. Creates a new game (or adds a new
// player to the existing game that corresponds with the provided gameID), and
// attempts to establish a Websocket connection with a client. Expects "gameID"
// as a query parameter when the /ws endpoint is hit.
func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract info from query params.
	queryParams := r.URL.Query()
	if _, ok := queryParams["gameID"]; !ok {
		errMsg := fmt.Sprint("the \"gameID\" query param is required")
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	gameID := queryParams["gameID"][0]

	// Look for an active game associated with the provided gameID. If one
	// doesn't exist, then create a new one.
	if _, ok := s.activeGames[gameID]; !ok {
		// TODO: Write function somewhere in model package to randomly generate
		// a dictionary. For now, I'm just using a dummy/stubbed one.
		dictionary := [25]string{"foo", "bar", "baz"}
		newGame := model.NewGame(dictionary)
		newInteractor := realtime.NewInteractor(newGame)
		s.activeGames[gameID] = newInteractor
	}

	upgrader := websocket.Upgrader{}
	// Set up upgrader to only accept connections from a few specified origins.
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return s.checkClientOrigin(r)
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
	newPlayer := model.NewPlayer(conn)
	s.activeGames[gameID].Players[newPlayer.DisplayName] = newPlayer
}

// checkClientOrigin makes sure that the client application that's trying to
// establish a websocket connection with us is someone we recognize: either the
// client's origin in development (localhost:some-port) or the client's
// production URL.
func (s *Server) checkClientOrigin(r *http.Request) bool {
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
