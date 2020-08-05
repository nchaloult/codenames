package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gorilla/mux"
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
		port,
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

// spaHandler implements the http.Handler interface, so we can use it to respond
// to HTTP requests. The path to the static directory and path to the index file
// within that static directory are used to serve the SPA in the given static
// directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir on the
// SPA handler. If a file is found, it will be served. If not, the file located
// at the index path on the SPA handler will be served. This is suitable
// behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating
		// the file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
