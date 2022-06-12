package forge_test

import (
	"testing"

	"github.com/fox091/forge"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("returns a random string with the default options", func(t *testing.T) {
		defaultCharSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		randStr := forge.String()
		assert.Len(t, randStr, 5)
		for _, char := range randStr {
			assert.Contains(t, defaultCharSet, string(char))
		}
	})

	t.Run("returns a random string of specified length", func(t *testing.T) {
		defaultCharSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		length := 20
		randStr := forge.String(func(so *forge.StringOptions) {
			so.Length = length
		})
		assert.Len(t, randStr, length)
		for _, char := range randStr {
			assert.Contains(t, defaultCharSet, string(char))
		}
	})

	t.Run("returns a random string using the specified character set", func(t *testing.T) {
		charSet := "abc"
		randStr := forge.String(func(so *forge.StringOptions) {
			so.CharSet = charSet
		})
		assert.Len(t, randStr, 5)
		for _, char := range randStr {
			assert.Contains(t, charSet, string(char))
		}
	})
}

func BenchmarkString(b *testing.B) {
	// Call it once to pre-cache the random source
	_ = forge.String()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		randStr := forge.String()
		assert.Len(b, randStr, 5)
	}
}
