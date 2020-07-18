package model

import "testing"

// TestGameCreation tests the NewGame() func, and that the primitive fields of
// the returned *Game object are initialized to their expected values.
func TestGameCreation(t *testing.T) {
	dictionary := [25]string{"word", "list"}
	got := NewGame(dictionary)

	// Make sure that all primitive fields are set appropriately for a Game that
	// has just begun.
	if got.RedTeamScore != 0 {
		t.Errorf("Unexpected RedTeamScore. got: %d, want: %d\n",
			got.RedTeamScore, 0)
	}
	if got.BlueTeamScore != 0 {
		t.Errorf("Unexpected BlueTeamScore. got: %d, want: %d\n",
			got.BlueTeamScore, 0)
	}
	if got.IsItRedTeamsTurn != true {
		t.Errorf("Unexpected value for IsItRedTeamsTurn. got: %v, want: %v\n",
			got.IsItRedTeamsTurn, true)
	}
	if got.RemainingFlips != -1 {
		t.Errorf("Unexpected value for RemainingFlips. got: %d, want: %d\n",
			got.RemainingFlips, -1)
	}
	if got.IsFinished != 0 {
		t.Errorf("Unexpected value for IsFinished. got: %d, want: %d\n",
			got.IsFinished, 0)
	}
}

// TestInvalidClues makes sure that a team's spymaster cannot give a clue with
// an invalid number: a number that's either greater than that team's number of
// remaining cards, or less than 1.
func TestInvalidClues(t *testing.T) {
	dictionary := [25]string{"foo", "bar"}
	game := NewGame(dictionary)

	tests := []struct {
		clue             int
		redTeamScore     int
		blueTeamScore    int
		isItRedTeamsTurn bool
	}{
		{-1, 0, 0, true},
		{0, 0, 0, true},
		{maxRedTeamScore, 1, 0, true},
		{maxBlueTeamScore, 0, 1, false},
	}
	for _, c := range tests {
		game.RedTeamScore = c.redTeamScore
		game.BlueTeamScore = c.blueTeamScore
		game.IsItRedTeamsTurn = c.isItRedTeamsTurn

		err := game.SetFlipsForCurrentTurn(c.clue)
		if err == nil {
			t.Fatalf("expected SetFlipsForCurrentTurn(%d) to return an error, but ti did not", c.clue)
		}
	}
}

// TestFlipOutOfBounds makes sure that the game state doesn't change if flipping
// a card that doesn't exist, or that is out of bounds, is attempted.
func TestFlipOutOfBounds(t *testing.T) {
	dictionary := [25]string{"foo", "bar"}
	game := NewGame(dictionary)

	tests := []struct {
		r int
		c int
	}{
		{-1, 0},
		{0, -1},
		{-1, -1},
		{boardSize, 0},
		{0, boardSize},
		{boardSize, boardSize},
	}
	for _, c := range tests {
		err := game.RevealCard(c.r, c.c)
		if err == nil {
			t.Fatalf("Expected RevealCard(%d, %d) to return an error, but it did not",
				c.r, c.c)
		}
	}
}

// TestCommonScenarios makes sure that the game score updates appropriately in a
// plethora of common situations or game states.
func TestCommonScenarios(t *testing.T) {
	dictionary := [25]string{"foo", "bar"}
	game := NewGame(dictionary)
	// Build "dummy" game board with Cards of all different category types.
	dummyGameBoard := gameBoard{
		[5]*Card{
			NewCard("neutral", 0),
			NewCard("neutral", 0),
			NewCard("redTeam", 1),
			NewCard("blueTeam", 2),
			NewCard("assassin", 3),
		},
	}
	game.Board = dummyGameBoard

	// Scenario: attempt to reveal a card before your team's spymaster has given
	// their clue for this turn.
	err := game.RevealCard(0, 0)
	if err == nil {
		t.Fatal("expected RevealCard(0, 0) to return error when RemainingFlips = -1")
	}

	game.SetFlipsForCurrentTurn(1)
	// Scenario: reveal a neutral card with your team's final flip this turn.
	err = game.RevealCard(0, 0)
	if err != nil {
		t.Fatalf("unexpected error when using team's final flip on a neutral card: %v",
			err)
	}
	if game.RedTeamScore != 0 {
		t.Fatalf("RedTeamScore changed after revealing a neutral card. got: %d, want: %d",
			game.RedTeamScore, 0)
	}
	if game.BlueTeamScore != 0 {
		t.Fatalf("BlueTeamScore changed after red team revealed a neutral card. got: %d, want: %d",
			game.BlueTeamScore, 0)
	}
	if game.IsItRedTeamsTurn {
		t.Fatal("team did not change after the use of a turn's final flip")
	}
	if game.RemainingFlips != -1 {
		t.Fatalf("unexpected RemainingFlips after a turn change. got: %d, want: %d",
			game.RemainingFlips, -1)
	}
	if game.IsFinished != 0 {
		t.Fatalf("unexpected IsFinished after a normal turn change. got: %d, want: %d",
			game.IsFinished, 0)
	}

	game.SetFlipsForCurrentTurn(2)
	// Scenario: reveal a neutral card with more flips remaining this turn.
	err = game.RevealCard(0, 1)
	if err != nil {
		t.Fatalf("unexpected error when revealing a neutral card with more flips remaining: %v",
			err)
	}
	if game.RedTeamScore != 0 {
		t.Fatalf("RedTeamScore changed after blue team revealed a neutral card. got: %d, want: %d",
			game.RedTeamScore, 0)
	}
	if game.BlueTeamScore != 0 {
		t.Fatalf("BlueTeamScore changed after revealing a neutral card. got: %d, want: %d",
			game.BlueTeamScore, 0)
	}
	if !game.IsItRedTeamsTurn {
		t.Fatal("team did not change after revealing a neutral card")
	}
	if game.RemainingFlips != -1 {
		t.Fatalf("unexpected RemainingFlips after a turn change. got: %d, want: %d",
			game.RemainingFlips, -1)
	}
	if game.IsFinished != 0 {
		t.Fatalf("unexpected IsFinished after a normal turn change. got: %d, want: %d",
			game.IsFinished, 0)
	}

	game.SetFlipsForCurrentTurn(1)
	// Scenario: reveal one of your team's cards with your team's final flip
	// this turn.
	err = game.RevealCard(0, 2)
	if err != nil {
		t.Fatalf("unexpected error when revealing one of your team's cards with your team's last flip this turn: %v",
			err)
	}
	if game.RedTeamScore != 1 {
		t.Fatalf("unexpected RedTeamScore after red team revealed one of their cards: got: %d, want: %d",
			game.RedTeamScore, 1)
	}
	if game.BlueTeamScore != 0 {
		t.Fatalf("BlueTeamScore changed after red team revealed one of their cards. got: %d, want: %d",
			game.BlueTeamScore, 0)
	}
	if game.IsItRedTeamsTurn {
		t.Fatal("team did not change after a last flip for this turn was used")
	}
	if game.RemainingFlips != -1 {
		t.Fatalf("unexpected RemainingFlips after a turn change. got: %d, want: %d",
			game.RemainingFlips, -1)
	}
	if game.IsFinished != 0 {
		t.Fatalf("unexpected IsFinished after a normal turn change. got: %d, want: %d",
			game.IsFinished, 0)
	}

	game.SetFlipsForCurrentTurn(2)
	// Scenario: reveal one of your team's cards with more flips remaining this
	// turn.
	err = game.RevealCard(0, 3)
	if err != nil {
		t.Fatalf("unexpected error when revealing one of your team's cards with more flips left: %v",
			err)
	}
	if game.RedTeamScore != 1 {
		t.Fatalf("unexpected RedTeamScore after blue team revealed one of their cards: got: %d, want: %d",
			game.RedTeamScore, 1)
	}
	if game.BlueTeamScore != 1 {
		t.Fatalf("unexpected BlueTeamScore after blue team revealed one of their cards: got: %d, want: %d",
			game.BlueTeamScore, 1)
	}
	if game.IsItRedTeamsTurn {
		t.Fatal("unexpected team change after a flip when more flips remain")
	}
	if game.RemainingFlips != 1 {
		t.Fatalf("unexpected RemainingFlips after a normal card flip. got: %d, want %d",
			game.RemainingFlips, 1)
	}
	if game.IsFinished != 0 {
		t.Fatalf("unexpected IsFinished after a normal card flip. got: %d, want: %d",
			game.IsFinished, 0)
	}

	// Scenario: reveal the assassin.
	err = game.RevealCard(0, 4)
	if err != nil {
		t.Fatalf("unexpected error when revealing the assassin: %v", err)
	}
	if game.RedTeamScore != 1 {
		t.Fatalf("unexpected RedTeamScore change after blue team revealed the assassin. got: %d, want: %d",
			game.RedTeamScore, 1)
	}
	if game.BlueTeamScore != 1 {
		t.Fatalf("unexpected BlueTeamScore change after blue team revealed the assassin. got: %d, want: %d",
			game.BlueTeamScore, 1)
	}
	if game.IsItRedTeamsTurn {
		t.Fatal("unexpected team change after the assassin was revealed")
	}
	if game.RemainingFlips != 1 {
		t.Fatal("unexpected RemainingFlips mutation after the assassin was revealed")
	}
	if game.IsFinished != 1 {
		t.Fatalf("unexpected IsFinished after blue team revealed the assassin. got: %d, want: %d",
			game.IsFinished, 1)
	}
}

// TestEndgameScenarios makes sure that the game state updates accordingly when
// a game-ending move is made.
func TestEndgameScenarios(t *testing.T) {
	dictionary := [25]string{"foo", "bar"}
	game := NewGame(dictionary)
	dummyGameBoard := gameBoard{
		[5]*Card{
			NewCard("redTeam", 1),
			NewCard("blueTeam", 2),
		},
	}

	tests := []struct {
		redTeamScore       int
		blueTeamScore      int
		isItRedTeamsTurn   bool
		expectedIsFinished int
	}{
		{maxRedTeamScore - 1, 0, true, 1},
		{0, maxBlueTeamScore - 1, false, 2},
	}
	for _, c := range tests {
		// Prepare the game state for the next test case.
		dummyGameBoard[0][0].IsRevealed = false
		dummyGameBoard[0][1].IsRevealed = false
		game.Board = dummyGameBoard
		game.RedTeamScore = c.redTeamScore
		game.BlueTeamScore = c.blueTeamScore
		game.IsItRedTeamsTurn = c.isItRedTeamsTurn
		game.IsFinished = 0

		game.SetFlipsForCurrentTurn(1)
		if c.isItRedTeamsTurn {
			game.RevealCard(0, 0)
		} else {
			game.RevealCard(0, 1)
		}

		if game.IsFinished != c.expectedIsFinished {
			t.Errorf("unexpected IsFinished: got: %d, want: %d",
				game.IsFinished, c.expectedIsFinished)
		}
	}
}
