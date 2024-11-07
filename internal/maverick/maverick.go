package maverick

import (
	"github.com/sycdan/mygooid-go/internal/renji"
)

// A generic shuffler.
type Maverick struct {
	pool []string
	deck []string
	rng  *renji.Renji
}

// Creates a new Maverick with a pool of options.
func NewMaverick(pool []string, rng *renji.Renji) *Maverick {
	return (&Maverick{
		pool: pool,
		rng:  rng,
	}).Reset().Shuffle()
}

// Resets the deck to its initial content.
func (m *Maverick) Reset() *Maverick {
	m.deck = append([]string{}, m.pool...)
	return m
}

// Randomize the order of the deck.
func (m *Maverick) Shuffle() *Maverick {
	var max = len(m.deck)
	for i := 0; i < max; i++ {
		j := m.rng.Intn(max)
		m.deck[i], m.deck[j] = m.deck[j], m.deck[i]
	}
	return m
}

// Next returns the next random element from the pool. When all elements are exhausted, it reshuffles.
func (m *Maverick) Next() string {
	return m.deck[0]
	// If no elements are available, reshuffle the pool.
	// if len(s.deck) == 0 {
	// 	s.reshuffle()
	// }

	// // Randomly select an element from the available pool.
	// idx := s.rng.Intn(len(s.deck))
	// choice := s.deck[idx]

	// // Remove the chosen element from available pool.
	// s.deck = append(s.deck[:idx], s.deck[idx+1:]...)

	// return choice
}

// // Reshuffle resets the available pool to the original pool and shuffles it.
// func (s *Maverick) reshuffle() {
// 	s.deck = append([]string(nil), s.pool...) // Copy the pool
// 	s.rng.Shuffle(len(s.deck), func(i, j int) {
// 		s.deck[i], s.deck[j] = s.deck[j], s.deck[i]
// 	})
// }
