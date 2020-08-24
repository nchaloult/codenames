package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthHandler handles requests to endpoints that respond with information
// about the server's current state and status.
type HealthHandler struct {
	server *Server
}

// NewHealthHandler returns a pointer to a new HealthHandler initialized with
// the provided fields.
func NewHealthHandler(server *Server) *HealthHandler {
	return &HealthHandler{server}
}

// healthHandler serves requests at the /health route. Responds with information
// about the server's state.
//
// TODO: include health information about the database in response.
func (h *HealthHandler) healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"ok":             true,
		"numActiveGames": len(h.server.activeGames),
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

// RegisterRoutes registers handlers for all of the routes that wsHandler
// supports.
func (h *HealthHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/health", h.healthHandler)
}
