package forge

var defaultWordList = []string{"hello", "world", "foo", "bar"}

type WordOptions struct {
	WordList []string
}

func Word(optionModifiers ...func(*WordOptions)) string {
	options := WordOptions{
		WordList: defaultWordList,
	}
	for _, opt := range optionModifiers {
		opt(&options)
	}
	wordListLength := len(options.WordList)

	idx := IndexGenerator.Index(wordListLength)
	return options.WordList[idx]
}
