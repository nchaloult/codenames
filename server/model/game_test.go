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
