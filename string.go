package forge

import (
	"math"
	"strings"
)

type StringOptions struct {
	Length  int
	CharSet string
}

const defaultStringLength = 5
const defaultCharSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func String(optionModifiers ...func(*StringOptions)) string {
	randSource := getRandSource()
	options := StringOptions{
		Length:  defaultStringLength,
		CharSet: defaultCharSet,
	}
	for _, opt := range optionModifiers {
		opt(&options)
	}
	requestedLength := options.Length
	charSetLength := len(options.CharSet)

	// Log2(n) will give the number of bits needed to store n when rounded up.
	// By adding 0.5, we ensure that `math.Round` will always round to the expected number,
	// since it rounds each direction from 0.5.
	var letterIndexBits = int(math.Round(math.Log2(float64(charSetLength)) + 0.5))

	// All 1-bits, as many as letterIndexBits
	var letterIndexMask = int64(1<<letterIndexBits - 1)

	// Number of letter indices fitting in 63 bits
	var letterIndexMax = int64(63 / letterIndexBits)

	sb := strings.Builder{}
	sb.Grow(requestedLength)

	for i, randBits, remainingIndexes := requestedLength-1, randSource.Int63(), letterIndexMax; i >= 0; {
		// If we run out of indexes, reset the random bits and remaining indexes.
		if remainingIndexes == 0 {
			randBits, remainingIndexes = randSource.Int63(), letterIndexMax
		}
		// Get an index from the lowest X bytes, but skip if it is too big to be an index.
		if idx := int(randBits & letterIndexMask); idx < charSetLength {
			sb.WriteByte(options.CharSet[idx])
			i--
		}
		// Shift out the used bits.
		randBits >>= letterIndexBits
		remainingIndexes--
	}

	return sb.String()
}
