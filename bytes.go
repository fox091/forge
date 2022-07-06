package forge

var byteMask = int64(1<<8 - 1)

// Byte returns a random byte.
func Byte() byte {
	rand := getRand()
	return byte(rand.Int63() & byteMask)
}

type BytesOptions struct {
	Count int
}

// Bytes returns a slice of random bytes.
//
// By default, it will give you 5 bytes, but you can specify the count you want.
func Bytes(optionModifiers ...func(*BytesOptions)) []byte {
	rand := getRand()
	options := BytesOptions{
		Count: 5,
	}
	for _, opt := range optionModifiers {
		opt(&options)
	}
	randBytes := make([]byte, options.Count)
	rand.Read(randBytes)
	return randBytes
}
