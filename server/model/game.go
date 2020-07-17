package model

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxRedTeamScore  = 9
	maxBlueTeamScore = 8
	boardSize        = 5
)

type gameBoard [boardSize][boardSize]*Card

// Game maintains state about an avtive game of Codenames, and exposes functions
// which mutate that state, allowing the game to progress.
type Game struct {
	Board            gameBoard
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
func NewGame(dictionary [25]string) *Game {
	return &Game{
		Board:            createBoard(dictionary),
		RedTeamScore:     0,
		BlueTeamScore:    0,
		IsItRedTeamsTurn: true,
		RemainingFlips:   -1,
		IsFinished:       0,
	}
}

// createBoard populates a gameBoard with 25 new pointers to Card objects, and
// returns it.
func createBoard(dictionary [25]string) gameBoard {
	rand.Seed(time.Now().UnixNano())

	// Shuffle the dictionary.
	rand.Shuffle(len(dictionary), func(i, j int) {
		dictionary[i], dictionary[j] = dictionary[j], dictionary[i]
	})

	// Create Cards with the appropriate classifications.
	cardNum := 0
	var board gameBoard
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			var classification int
			if cardNum < maxRedTeamScore {
				classification = 1
			} else if cardNum < maxRedTeamScore+maxBlueTeamScore {
				classification = 2
			} else if cardNum < maxRedTeamScore+maxBlueTeamScore+1 {
				classification = 3
			} else {
				classification = 0
			}

			board[r][c] = NewCard(dictionary[cardNum], classification)
			cardNum++
		}
	}

	// Shuffle the board after assigning Cards their classifications in a linear
	// fashion. Inspired by the Fisher-Yates algorithm.
	for i := len(board) - 1; i > 0; i-- {
		for j := len(board[i]) - 1; j > 0; j-- {
			m := rand.Intn(i + 1)
			n := rand.Intn(j + 1)

			temp := board[i][j]
			board[i][j] = board[m][n]
			board[m][n] = temp
		}
	}

	return board
}

// ChangeTurn changes which team's turn it is.
func (g *Game) ChangeTurn() {
	g.IsItRedTeamsTurn = !g.IsItRedTeamsTurn
	g.RemainingFlips = -1
}

// RevealCard reveals the Card on the rth row and the cth column of the
// gameBoard. Calls that Card's Reveal() method, and updates the game state
// accordingly.
func (g *Game) RevealCard(r, c int) error {
	// Input validation.
	if r < 0 || r >= boardSize || c < 0 || c >= boardSize {
		return fmt.Errorf("The provided r (%d) and c (%d) must be within the range: [0, %d)",
			r, c, boardSize)
	}

	if g.IsFinished != 0 {
		return fmt.Errorf("cannot reveal a card after the game has finished")
	}
	if g.RemainingFlips <= 0 {
		return fmt.Errorf("cannot reveal a card when RemainingFlips = %d",
			g.RemainingFlips)
	}

	class := g.Board[r][c].Reveal()
	if class == -1 {
		// The card has already been revealed. NOP.
		return nil
	} else if class == 0 {
		// The card is neutral. End the turn without touching the game's score.
		g.ChangeTurn()
		return nil
	} else if class == 1 {
		g.RedTeamScore++
	} else if class == 2 {
		g.BlueTeamScore++
	} else if class == 3 {
		// The card is the assassin. End the game immediately.
		if g.IsItRedTeamsTurn {
			g.IsFinished = 2
		} else {
			g.IsFinished = 1
		}
		return nil
	}

	g.IsFinished = g.EvaluateWinner()
	if g.IsFinished > 0 {
		return nil
	}

	g.RemainingFlips--
	if g.RemainingFlips <= 0 {
		g.ChangeTurn()
	}

	return nil
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
