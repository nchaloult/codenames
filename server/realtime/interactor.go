package realtime

import (
	"github.com/nchaloult/codenames/model"
)

// Interactor manages Players that are playing in a Game, and facilitates events
// through Websocket connections with all of those Players. When a Player
// performs an action that triggers an event, whether that action be in an
// ongoing game or a game lobby, an Interactor processes that event by mutating
// the data model and notifying all other Players.
//
// Interactor and Game have a 1-1 relationship; there is one Interactor instance
// associated with, or responsible for, each Game.
type Interactor struct {
	Game *model.Game

	// Players stores Player objects for each active client, indexed by their
	// display name.
	Players map[string]*Player
}

// NewInteractor returns a pointer to a new Interactor object initialized with
// the provided game.
func NewInteractor(game *model.Game) *Interactor {
	return &Interactor{
		Game:    game,
		Players: make(map[string]*Player, 0),
	}
}
