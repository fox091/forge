package forge

type BoolOptions struct {
	PercentTrue float64
}

// Bool returns true or false randomly.
//
// The default options will return true 50% of the time.
//
// PercentTrue can be used to skew the chance of receiving true.
// For example: PercentTrue: 0.3 will return true 30% of the time.
// 1.0 will return true 100% of the time.
func Bool(optionModifiers ...func(*BoolOptions)) bool {
	rand := getRand()
	options := BoolOptions{
		PercentTrue: 0.5,
	}
	for _, opt := range optionModifiers {
		opt(&options)
	}

	return rand.Float64() < options.PercentTrue
}
