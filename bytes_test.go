package forge_test

import (
	"testing"

	"github.com/fox091/forge"
	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	t.Run("returns 5 bytes by default", func(t *testing.T) {
		randBytes := forge.Bytes()
		assert.Len(t, randBytes, 5)
	})

	t.Run("returns the specified number of random bytes", func(t *testing.T) {
		count := 200
		optionModifier := func(bo *forge.BytesOptions) {
			bo.Count = count
		}
		randBytes := forge.Bytes(optionModifier)
		assert.Len(t, randBytes, count)
	})
}
