package realtime

import (
	"github.com/google/uuid"
	"github.com/nchaloult/codenames/model"
)

// Interactor facilitates events through a Websocket connection with clients.
// Each Interactor instance is associated with a Game instance; they have a 1-1
// relationship. Whenever a client performs an action that triggers an event,
// an Interactor instance processes that
// Players/connected clients.

// Interactor manages Players that are playing in a Game, and facilitates events
// through Websocket connections with all of those Players. When a Player
// performs an action that triggers an event, whether that action be in an
// ongoing game or a game lobby, an Interactor processes that event by mutating
// the data model and notifying all other Players.
//
// Interactor and Game have a 1-1 relationship; there is one Interactor instance
// associated with, or responsible for, each Game.
type Interactor struct {
	Game    *model.Game
	Players map[uuid.UUID]model.Player
}

// NewInteractor returns a pointer to a new Interactor object initialized with
// the provided game.
func NewInteractor(game *model.Game) *Interactor {
	return &Interactor{
		Game:    game,
		Players: make(map[uuid.UUID]model.Player, 0),
	}
}
