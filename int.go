package forge

import (
	"fmt"
	"math"
	"math/bits"
	"math/rand"
)

type IntOptions struct {
	Min int
	Max int
}

// Int returns a random number in the range [min, max).
//
// The default options are Min: math.MinInt, Max: math.MaxInt.
//
// If the minimum is equal to the maximum, it will return the maximum
//
// The calculations are platform independent and will step down to 32 bit functions on a 32 bit architecture.
func Int(optionModifiers ...func(*IntOptions)) (int, error) {
	seed()
	options := IntOptions{
		Min: math.MinInt,
		Max: math.MaxInt,
	}
	for _, modifier := range optionModifiers {
		modifier(&options)
	}
	if options.Min == options.Max {
		return options.Max, nil
	}
	if options.Min > options.Max {
		return 0, fmt.Errorf("invalid options: Min (%d) cannot be greater than Max (%d)", options.Min, options.Max)
	}
	min := uint(options.Min)
	var chosenNumber uint
	// Use float32s if 32 bit architecture.  Likely not needed, but keeps full 32/64 bit compatibility.
	if bits.UintSize == 32 {
		// Get the max number desired as if the minimum was 0
		maxZeroBasedDesiredNumber := float32(uint(options.Max) - min)
		// This will give us a number between 0 and our 0-based max
		chosenNumber = uint(rand.Float32() * maxZeroBasedDesiredNumber)
	} else {
		// Get the max number desired as if the minimum was 0
		maxZeroBasedDesiredNumber := float64(uint(options.Max) - min)
		// This will give us a number between 0 and our 0-based max
		chosenNumber = uint(rand.Float64() * maxZeroBasedDesiredNumber)
	}
	// Add the minimum back to go from 0-based back to the range we really want.
	return int(chosenNumber + min), nil
}
