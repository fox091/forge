package forge_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/fox091/forge"
	"github.com/stretchr/testify/assert"
)

func Test_Int(t *testing.T) {
	for i := 0; i < 10000; i++ {
		res, err := forge.Int()
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, res, math.MinInt)
		assert.Less(t, res, math.MaxInt)
		fmt.Println(res)
	}
}
