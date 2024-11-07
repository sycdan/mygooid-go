package renji

import (
	"encoding/hex"
	"math"

	"github.com/sycdan/mygooid-go/internal/utils"
)

type Renji struct {
	lastHash *string
	seedHash string
}

// Create a new instance with the given seed hash.
func NewRenji(seedHash string) *Renji {
	return (&Renji{seedHash: seedHash}).Reset()
}

// Reset the last hash to the seed hash.
func (r *Renji) Reset() *Renji {
	r.lastHash = &r.seedHash
	return r
}

// Compute and return the next hash in the sequence.
func (r *Renji) nextHash() string {
	newHash := utils.HashText(*r.lastHash)
	r.lastHash = &newHash
	return newHash
}

// Generates a number between 0 and 1 (inclusive).
func (r *Renji) Float64() float64 {
	hash := r.nextHash()

	bytes, err := hex.DecodeString(hash[:64])
	if err != nil {
		utils.Die(err.Error())
	}

	var bigEight uint64 // Positive values only
	for i := 0; i < 8; i++ {
		bigEight |= uint64(bytes[i]) << (56 - 8*i)
	}

	return float64(bigEight) / float64(math.MaxUint64)
}

// Generates a random integer between 0 (inclusive) and `lessThan` (exclusive).
func (r *Renji) Intn(lessThan int) int {
	return int(math.Floor(r.Float64() * float64(lessThan)))
}

// Return a random integer between `min` and `max` (inclusive).
func (r *Renji) NextBetween(min int, max int) int {
	return min + r.Intn(max-min+1)
}
