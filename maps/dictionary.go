package maps

type Dictionary map[string]string

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrNotFound        = DictionaryError("could not find the word you were looking for")
	ErrWordExists      = DictionaryError("cannot add word because it already exists")
	ErrorWordNotExists = DictionaryError("cannot update word because it is not exists")
)

func (d Dictionary) Search(word string) (string, error) {
	if v, ok := d[word]; ok {
		return v, nil
	} else {
		return "", ErrNotFound
	}
}

func (d Dictionary) Add(word, definition string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	} else {
		d[word] = definition
		return nil
	}
}

func (d Dictionary) Update(word, definition string) error {
	if _, ok := d[word]; ok {
		d[word] = definition
		return nil
	} else {
		d[word] = definition
		return ErrorWordNotExists
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
