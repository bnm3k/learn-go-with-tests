package c8

type Dictionary map[string]string

const (
	ErrNotFound               = DictionaryErr("cannot not find the key in dict")
	ErrWordExists             = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist       = DictionaryErr("cannot  update an entry whose key not present in dict")
	ErrDeleteWordDoesNotExist = DictionaryErr("cannot key an entry whose key not present in dict")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
func (dict Dictionary) Search(word string) (string, error) {
	entry, ok := dict[word]
	if !ok {
		return "", ErrNotFound
	}
	return entry, nil
}

func (dict Dictionary) Add(key, entry string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		dict[key] = entry
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Update(key, newEntry string) error {
	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dict[key] = newEntry
	default:
		return err
	}

	return nil
}

func (dict Dictionary) Delete(key string) error {

	_, err := dict.Search(key)

	switch err {
	case ErrNotFound:
		return ErrDeleteWordDoesNotExist
	case nil:
		delete(dict, key)
		return nil
	default:
		return err
	}

}
