package dictionary

const (
	NoDefinitionError     = DictionaryErr("no definition present for word")
	WordExistsError       = DictionaryErr("word already exists")
	WordDoesNotExistError = DictionaryErr("word does not exist to update")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary struct {
	dictionary map[string]string
}

func (d *Dictionary) Add(word string, definition string) error {
	if d.dictionary == nil {
		d.dictionary = map[string]string{word: definition}
		return nil
	}
	storedDefinition, err := d.Search(word)

	if err == NoDefinitionError {
		d.dictionary[word] = definition
		return nil
	}

	if storedDefinition == definition {
		return nil
	}

	return WordExistsError
}

func (d *Dictionary) Search(word string) (string, error) {
	definition := d.dictionary[word]
	if definition != "" {
		return definition, nil
	}
	return "", NoDefinitionError
}

func (d *Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	if err == NoDefinitionError {
		return WordDoesNotExistError
	}
	d.dictionary[word] = definition
	return nil
}

func (d *Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == NoDefinitionError {
		return WordDoesNotExistError
	}
	delete(d.dictionary, word)
	return nil
}
