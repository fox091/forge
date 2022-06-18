package forge

import (
	"math/rand"
	"time"
)

var cachedRandSource rand.Source

func getRandSource() rand.Source {
	if cachedRandSource == nil {
		cachedRandSource = rand.NewSource(time.Now().UnixNano())
	}
	return cachedRandSource
}

var cachedRand *rand.Rand

func getRand() *rand.Rand {
	if cachedRand == nil {
		cachedRand = rand.New(getRandSource())
	}
	return cachedRand
}
