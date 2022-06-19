package forge

import (
	"strings"
)

type StringOptions struct {
	Length  int
	CharSet string
}

const defaultStringLength = 5
const defaultCharSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// String returns a randomly generated string.
//
// By default, it will give you a string with a length of 5.
//
// The default character set used is all of the lowercase and uppercase english letters.
//
// A custom character set, and length can be provided.
func String(optionModifiers ...func(*StringOptions)) string {
	options := StringOptions{
		Length:  defaultStringLength,
		CharSet: defaultCharSet,
	}
	for _, opt := range optionModifiers {
		opt(&options)
	}
	requestedLength := options.Length
	charSetLength := len(options.CharSet)

	sb := strings.Builder{}
	sb.Grow(requestedLength)

	for i := 0; i < requestedLength; i++ {
		idx := IndexGenerator.Index(charSetLength)
		sb.WriteByte(options.CharSet[idx])
	}

	return sb.String()
}
