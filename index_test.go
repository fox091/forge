package forge_test

import (
	"testing"

	"github.com/fox091/forge"
	"github.com/stretchr/testify/assert"
)

func TestIndexGenerator(t *testing.T) {
	indexMap := map[int]float64{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}
	iterations := 1000000
	numIndexes := 10
	for i := 0; i < iterations; i++ {
		index = forge.IndexGenerator.Index(numIndexes)
		assert.Less(t, index, numIndexes)
		assert.GreaterOrEqual(t, index, 0)
		indexMap[index] += 1
	}
	targetIndexCount := iterations / numIndexes
	for _, count := range indexMap {
		assert.Less(t, count, float64(targetIndexCount)*1.05)
		assert.Greater(t, count, float64(targetIndexCount)*0.95)
	}
}

var index int

func BenchmarkIndexGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		index = forge.IndexGenerator.Index(100)
	}
}
