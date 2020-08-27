package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nchaloult/codenames/realtime"
)

// HealthHandler handles requests to endpoints that respond with information
// about the server's current state and status.
type HealthHandler struct {
	manager *realtime.Manager
}

// NewHealthHandler returns a pointer to a new HealthHandler initialized with
// the provided fields.
func NewHealthHandler(manager *realtime.Manager) *HealthHandler {
	return &HealthHandler{manager}
}

// healthHandler serves requests at the /health route. Responds with information
// about the server's state.
//
// TODO: include health information about the database in response.
func (h *HealthHandler) healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"ok":             true,
		"numActiveGames": len(h.manager.ActiveGames),
	}
	constructAndSendResponse(w, response)
}

// RegisterRoutes registers handlers for all of the routes that wsHandler
// supports.
func (h *HealthHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/health", h.healthHandler)
}
