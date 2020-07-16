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
