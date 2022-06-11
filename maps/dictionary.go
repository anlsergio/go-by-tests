package maps

type Dictionary map[string]string

var (
	ErrNotFound   = DictonaryErr("could not find the word you were looking for")
	ErrWordExists = DictonaryErr("the word you are trying to add already exists")
)

type DictonaryErr string

func (d DictonaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
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
