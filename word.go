package forge

var defaultWordList = []string{"hello", "world", "foo", "bar"}

type WordOptions struct {
	WordList []string
}

// Word returns a random word from a list of words.
//
// It uses a default simple word list and you can provide a custom word list to use instead.
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
