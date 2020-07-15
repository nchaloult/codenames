package model

// Card stores information about one of the 25 cards on a Codenames board.
type Card struct {
	Word       string
	IsRevealed bool

	// Classification represents what type of card this is. Is set to 0 if this
	// card is a neutral card, 1 if it's one of red team's cards, 2 if it's one
	// of blue team's cards, and 3 if it's the assassin.
	Classification int
}

// NewCard returns a pointer to a new Card object that's configured with the
// provided field values.
func NewCard(word string, classification int) *Card {
	return &Card{
		Word:           word,
		IsRevealed:     false,
		Classification: classification,
	}
}

// Reveal sets this Card's IsRevealed field to true. Returns this Card's
// classification.
//
// A Card that's been revealed can't be "un"-revealed; revealing a Card is an
// idempotent operation.
func (c *Card) Reveal() int {
	c.IsRevealed = true
	return c.Classification
}
