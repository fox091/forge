package forge

import (
	"math"
	"math/rand"
)

var IndexGenerator = NewRandomIndexGenerator()

type RandomIndexGenerator struct {
	randSource    rand.Source
	randBits      int64
	indexBits     int
	indexMask     int64
	remainingBits int
	itemCount     int
}

func NewRandomIndexGenerator() RandomIndexGenerator {
	randSource := getRandSource()

	randBits := randSource.Int63()

	return RandomIndexGenerator{
		randSource:    randSource,
		randBits:      randBits,
		remainingBits: 63,
	}
}

func (g *RandomIndexGenerator) Index(itemCount int) int {
	if itemCount != g.itemCount {
		// Log2(n) will give the number of bits needed to represent n when rounded up.
		// By adding 0.5, we ensure that `math.Round` will always round to the expected number,
		// since it rounds each direction from 0.5.
		g.indexBits = int(math.Round(math.Log2(float64(itemCount)) + 0.5))

		// All 1-bits, as many as indexBits
		g.indexMask = int64(1<<g.indexBits - 1)
		g.itemCount = itemCount
	}
	for {
		// If we run out of indexes, reset the random bits and remaining indexes.
		if g.remainingBits < g.indexBits {
			g.remainingBits = 63
			g.randBits = g.randSource.Int63()
		}

		// Get an index from the lowest X bytes
		idx := int(g.randBits & g.indexMask)

		// Shift out the used bits.
		g.randBits >>= g.indexBits
		g.remainingBits -= g.indexBits

		// Only return the index if it's small enough to be an index for this number of items.
		if idx < g.itemCount {
			return idx
		}
	}
}
