package forge

import (
	"math/rand"
	"time"
)

var hasSeeded = false

func seed() {
	if hasSeeded {
		return
	}
	rand.Seed(time.Now().UnixNano())
}
