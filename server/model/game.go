package model

import "fmt"

const (
	maxRedTeamScore  = 9
	maxBlueTeamScore = 8
)

// Game maintains state about an avtive game of Codenames, and exposes functions
// which mutate that state, allowing the game to progress.
type Game struct {
	RedTeamScore     int
	BlueTeamScore    int
	IsItRedTeamsTurn bool

	// RemainingFlips is set to -1 if a team has ended their turn, and the other
	// team's spymaster has yet to deliver their clue.
	RemainingFlips int

	// IsFinished stores whether this game has concluded or not. Is set to 0 if
	// the game isn't finished yet, 1 if the red team won, or 2 if the blue team
	// won.
	IsFinished int
}

// NewGame returns a pointer to a new Game object with initialized fields.
func NewGame() (*Game, error) {
	// TODO: accept dictionary []string as an arg, and call createBoard()
	return &Game{
		RedTeamScore:     0,
		BlueTeamScore:    0,
		IsItRedTeamsTurn: true,
		RemainingFlips:   -1,
		IsFinished:       0,
	}, nil
}

// ChangeTurn changes which team's turn it is.
func (g *Game) ChangeTurn() {
	g.IsItRedTeamsTurn = !g.IsItRedTeamsTurn
	g.RemainingFlips = -1
}

// SetFlipsForCurrentTurn sets the number of remaining card flips for the active
// turn. Typically called after ChangeTurn() is called and the active team's
// spymaster delivers their clue.
func (g *Game) SetFlipsForCurrentTurn(flips int) error {
	// Input validation.
	if flips <= 0 {
		return fmt.Errorf("flips: %d must be greater than 0", flips)
	}
	redTeamNumUnrevealedCards := maxRedTeamScore - g.RedTeamScore
	if g.IsItRedTeamsTurn && flips > redTeamNumUnrevealedCards {
		return fmt.Errorf("flips: %d must not be greater than red team's"+
			" number of remaining unrevealed cards: %d",
			flips,
			redTeamNumUnrevealedCards,
		)
	}
	blueTeamNumUnrevealedCards := maxBlueTeamScore - g.BlueTeamScore
	if !g.IsItRedTeamsTurn && flips > blueTeamNumUnrevealedCards {
		return fmt.Errorf("flips: %d must not be greater than blue team's"+
			" number of remaining unrevealed cards: %d",
			flips,
			blueTeamNumUnrevealedCards,
		)
	}

	g.RemainingFlips = flips
	return nil
}

// EvaluateWinner re-evaluates the state of the game, mutating the game's
// IsFinished field, if necessary. Returns whether the game is still ongoing,
// the red team won, or the blue team won.
func (g *Game) EvaluateWinner() int {
	// Assumes that both teams can't both reach their max scores.
	if g.RedTeamScore >= maxRedTeamScore {
		g.IsFinished = 1
	} else if g.BlueTeamScore >= maxBlueTeamScore {
		g.IsFinished = 2
	}
	return g.IsFinished
}
