package forge

type BytesOptions struct {
	Count int
}

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
