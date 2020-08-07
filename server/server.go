package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/nchaloult/codenames/realtime"
)

const (
	minPort = 1025
	maxPort = 65535
)

// Server exposes HTTP API endpoints that let clients mutate and interact with
// a Game's state, serves the static files for the web front-end, and maintains
// a collection of active Games in memory.
type Server struct {
	// port is the port number that an HTTP server will listen for requests on.
	port int

	// activeGames stores pointers to Interactors for all ongoing games that the
	// server is handling.
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
	router.HandleFunc("/api/hello", s.helloHandler).Methods("GET")

	// Stand up the server.
	log.Printf("Listening on port %d....\n", s.port)
	portAddr := fmt.Sprintf(":%d", s.port)
	log.Fatal(http.ListenAndServe(portAddr, router))
}

// helloHandler serves requests at the /api/hello route. Responds with a
// greeting as plain text.
func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}
