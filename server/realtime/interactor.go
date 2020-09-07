package realtime

import (
	"github.com/nchaloult/codenames/model"
)

// Interactor manages Players that are playing in a Game. When a Player triggers
// an event that all other Players in a game need to know about, an Interactor
// notifies those other Players by sending an event to each of them.
//
// Interactor and Game have a 1-1 relationship; there is one Interactor instance
// associated with, or responsible for, each Game.
type Interactor struct {
	Game *model.Game
	// Players stores Player objects for each active client, indexed by their
	// IDs.
	Players map[string]*Player
	// Events from a Player that are to be broadcasted to all other Players in a
	// Game.
	msgsToBroadcast chan *broadcastMsg
}

// NewInteractor returns a pointer to a new Interactor object initialized with
// the provided game.
func NewInteractor(game *model.Game) *Interactor {
	return &Interactor{
		Game:            game,
		Players:         make(map[string]*Player, 0),
		msgsToBroadcast: make(chan *broadcastMsg),
	}
}

// ListenForBroadcasts watches the msgsToBroadcast channel for any new events.
// When a new event appears, broadcast that event to all other Players in a Game
// except for the Player who originally created that event.
func (i *Interactor) ListenForBroadcasts() {
	for {
		select {
		case msg := <-i.msgsToBroadcast:
			for _, player := range i.Players {
				if player.ID == msg.originClientID {
					continue
				}
				select {
				case player.broadcastedMsgs <- msg.event:
					// No-op. Just push the broadcasted event onto each Player's
					// chan.
				default:
					// TODO: better error handling. For now, just assume that
					// this Player disconnected or something lol.
					close(player.broadcastedMsgs)
					delete(i.Players, player.ID)
				}
			}
		}
	}
}
