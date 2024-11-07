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
func New(seedHash string) *Renji {
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

// Convert the next hash to a number between 0 and 1.
func (r *Renji) Next() float64 {
	hash := r.nextHash()

	bytes, err := hex.DecodeString(hash[:64])
	if err != nil {
		utils.Die(err.Error())
	}

	var bigEight uint64
	for i := 0; i < 8; i++ {
		bigEight |= uint64(bytes[i]) << (56 - 8*i)
	}

	return float64(bigEight) / float64(math.MaxUint64)
}
