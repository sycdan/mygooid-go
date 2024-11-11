package table

import "github.com/sycdan/mygooid-go/internal/maverick"

type Table struct {
	Name  string
	Decks []maverick.Maverick
}

func NewTable(anchorCharacters string) *Table {
	t := &Table{Name: anchorCharacters, Decks: make([]maverick.Maverick, 0)}
	return t
}

func (t *Table) AddDeck(deck *maverick.Maverick) {
	t.Decks = append(t.Decks, *deck)
}
