package model

// Player stores information about a player in a Game. Players are managed by
// an Interactor.
type Player struct {
	DisplayName string
	IsOnRedTeam bool
	IsSpymaster bool
}

// NewPlayer returns a pointer to a new Player object with initialized fields.
func NewPlayer(displayName string, isOnRedTeam, isSpymaster bool) *Player {
	return &Player{displayName, isOnRedTeam, isSpymaster}
}
