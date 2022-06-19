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

	for i := 0; i < requestedLength; i += 1 {
		idx := IndexGenerator.Index(charSetLength)
		sb.WriteByte(options.CharSet[idx])
	}

	return sb.String()
}
