package handlers

import (
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

// Handler sets up handlers for HTTP endpoints, and spins up an HTTP server.
type Handler struct {
	router *mux.Router
	port   int
}

// NewHandler returns a pointer to a new Handler object initialized with the
// provided pointer to router.
func NewHandler(router *mux.Router, port int) (*Handler, error) {
	if port < minPort || port > maxPort {
		return nil, fmt.Errorf(
			"the provided port must be within the range: [%d, %d]",
			minPort,
			maxPort,
		)
	}

	return &Handler{router, port}, nil
}

// ListenOnEndpoints registers all of the HTTP handler funcs with their
// corresponding routes and begins listening for new requests that come in on
// those routes.
func (h *Handler) ListenOnEndpoints(handlers []RouteHandler) {
	// Register routes with their corresponding handler funcs.
	for _, handler := range handlers {
		handler.RegisterRoutes(h.router)
	}

	// Start accepting requests on those routes.
	log.Printf("Listening on port %d....\n", h.port)
	portAddr := fmt.Sprintf(":%d", h.port)
	log.Fatal(http.ListenAndServe(portAddr, h.router))
}
