package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/nchaloult/codenames/realtime"
	"github.com/nchaloult/codenames/server/handlers"
)

// Server exposes HTTP API endpoints that let clients mutate and interact with
// a Game's state, serves the static files for the web front-end, and maintains
// a collection of active Games in memory.
type Server struct {
	// activeGames stores pointers to Interactors for all ongoing games that the
	// server is handling, indexed by gameIDs.
	activeGames map[string]*realtime.Interactor
}

// NewServer returns a pointer to a new Server object that's configured with the
// provided port number.
func NewServer() *Server {
	return &Server{
		activeGames: make(map[string]*realtime.Interactor, 0),
	}
}

// Start spins up all of the Server's goroutines and brings everything online.
func (s *Server) Start(port int) error {
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

	// Set up HTTP handlers.
	router := mux.NewRouter()
	handler, err := handlers.NewHandler(router, port)
	if err != nil {
		return err
	}
	go handler.ListenOnEndpoints([]handlers.RouteHandler{
		NewHealthHandler(s),
		NewWSHandler(s),
	})

	return nil
}
