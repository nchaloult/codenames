package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	minPort = 1025
	maxPort = 65535
)

// RouteHandler describes objects which handle requests to specific HTTP
// endpoints.
type RouteHandler interface {
	RegisterRoutes(*mux.Router)
}

// API sets up handlers for HTTP endpoints, and spins up an HTTP server.
type API struct {
	router *mux.Router
	port   int
}

// NewAPI returns a pointer to a new API object initialized with the provided
// pointer to router and port number to listen on.
func NewAPI(router *mux.Router, port int) (*API, error) {
	if port < minPort || port > maxPort {
		return nil, fmt.Errorf(
			"the provided port must be within the range: [%d, %d]",
			minPort,
			maxPort,
		)
	}

	return &API{router, port}, nil
}

// ListenOnEndpoints registers all of the HTTP handler funcs with their
// corresponding routes and begins listening for new requests that come in on
// those routes.
func (a *API) ListenOnEndpoints(handlers []RouteHandler) {
	// Register routes with their corresponding handler funcs.
	for _, handler := range handlers {
		handler.RegisterRoutes(a.router)
	}

	// Start accepting requests on those routes.
	log.Printf("Listening on port %d....\n", a.port)
	portAddr := fmt.Sprintf(":%d", a.port)
	log.Fatal(http.ListenAndServe(portAddr, a.router))
}

// constructAndSendResponse adds important, common headers to endpoint
// responses, and marshals the provided response body into JSON.
func constructAndSendResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		errMsg := fmt.Sprintf("failed to encode response as JSON: %v", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
}
