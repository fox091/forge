package forge_test

import (
	"testing"

	"github.com/fox091/forge"
	"github.com/stretchr/testify/assert"
)

var defaultWordList = []string{"hello", "world", "foo", "bar"}

func TestWord(t *testing.T) {
	t.Run("returns a random word from the default word list", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			word := forge.Word()
			assert.Contains(t, defaultWordList, word)
		}
	})

	t.Run("returns a random word from a custom word list", func(t *testing.T) {
		customWordList := []string{"test", "words", "here"}
		optionModifier := func(wo *forge.WordOptions) {
			wo.WordList = customWordList
		}
		for i := 0; i < 1000; i++ {
			word := forge.Word(optionModifier)
			assert.Contains(t, customWordList, word)
		}
	})
}
