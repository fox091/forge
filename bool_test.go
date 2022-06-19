package forge_test

import (
	"testing"

	"github.com/fox091/forge"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	t.Run("returns true 50 percent of the time by default", func(t *testing.T) {
		trueCount, falseCount := 0, 0
		for i := 0; i < 100000; i++ {
			randBool := forge.Bool()
			if randBool {
				trueCount += 1
			} else {
				falseCount += 1
			}
		}
		assert.Greater(t, float64(trueCount)/float64(falseCount), 0.95)
	})

	t.Run("returns true a custom percent of the time", func(t *testing.T) {
		boolOptionFunc := func(bo *forge.BoolOptions) {
			bo.PercentTrue = 0.1
		}
		trueCount, falseCount := 0, 0
		for i := 0; i < 100000; i++ {
			randBool := forge.Bool(boolOptionFunc)
			if randBool {
				trueCount += 1
			} else {
				falseCount += 1
			}
		}
		trueToFalseRatio := float64(trueCount) / float64(falseCount)
		assert.Greater(t, trueToFalseRatio, 0.105555)
		assert.Less(t, trueToFalseRatio, 0.115555)
	})
}
