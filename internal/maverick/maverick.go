package maverick

import (
	"github.com/sycdan/mygooid-go/internal/renji"
)

// A generic shuffler.
type Maverick struct {
	pool []interface{}
	deck []interface{}
	rng  *renji.Renji
}

type MaverickPSlice []*Maverick

// Convert a typed slice to an interface slice.
func (m MaverickPSlice) ToISlice() []interface{} {
	var iSlice []interface{}
	for _, v := range m {
		iSlice = append(iSlice, v)
	}
	return iSlice
}

// Creates a new Maverick with a pool of options.
func NewMaverick(pool []interface{}, rng *renji.Renji) *Maverick {
	return (&Maverick{
		pool: pool,
		rng:  rng,
	}).Reset().Shuffle()
}

// Resets the deck to its initial content.
func (m *Maverick) Reset() *Maverick {
	m.deck = append([]interface{}{}, m.pool...)
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

// Returns the next element from the pool. When all elements are exhausted, it reshuffles.
func (m *Maverick) Next() interface{} {
	if len(m.deck) == 0 {
		m.Reset()
		m.Shuffle()
	}

	// Get the top card from the deck.
	topCard := m.deck[0]

	// Remove the top card from the deck.
	m.deck = m.deck[1:]

	return topCard
}

func Draw[T any](m *Maverick) T {
	return (m.Next()).(T)
}
