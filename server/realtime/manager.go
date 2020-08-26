package realtime

// Manager keeps track of all active games.
type Manager struct {
	// ActiveGames stores pointers to Interactors for all ongoing games that the
	// server is handling, indexed by gameIDs.
	ActiveGames map[string]*Interactor
}

// NewManager returns a pointer to a new Manager object with initialized fields.
func NewManager() *Manager {
	return &Manager{
		ActiveGames: make(map[string]*Interactor, 0),
	}
}
