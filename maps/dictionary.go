package maps

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("the word you are trying to add already exists")
	ErrWordDoesNotExist = DictionaryErr("the word you are trying to update doesn't exist")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err != ErrNotFound {
		if err != nil {
			return err
		}

		return ErrWordExists
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		if err == ErrNotFound {
			return ErrWordDoesNotExist
		}
		return err
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
