package forge_test

import (
	"testing"

	"github.com/fox091/forge"
	"github.com/stretchr/testify/assert"
)

func TestByte(t *testing.T) {
	t.Run("returns an relatively even distribution of random bytes", func(t *testing.T) {
		byteMap := map[byte]float64{}
		iterations := 10000000
		numBytes := 256
		for i := 0; i < iterations; i++ {
			b := forge.Byte()
			byteMap[b] += 1
		}
		targetByteCount := iterations / numBytes
		for _, count := range byteMap {
			assert.Less(t, count, float64(targetByteCount)*1.05)
			assert.Greater(t, count, float64(targetByteCount)*0.95)
		}
	})
}

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
