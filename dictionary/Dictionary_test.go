package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	dictionary := &Dictionary{}
	err := dictionary.Add("test", "this is just a test definition")
	assert.Nil(t, err)

	definition, err := dictionary.Search("test")
	assert.Nil(t, err)
	expectedDefinition := "this is just a test definition"

	assert.Equal(t, expectedDefinition, definition)
}

func TestAddTwoWords(t *testing.T) {
	dictionary := &Dictionary{}
	err := dictionary.Add("test", "this is just a test definition")
	assert.Nil(t, err)
	err = dictionary.Add("other", "this is just a different word")
	assert.Nil(t, err)

	assert.Equal(t, 2, len(dictionary.dictionary))
}

func TestUnknownWord(t *testing.T) {
	dictionary := &Dictionary{}
	dictionary.Add("test", "this is just a test definition")

	definition, err := dictionary.Search("not present")

	assert.Equal(t, "", definition)
	assert.Equal(t, NoDefinitionError, err)
}

func TestAddWordAlreadyExists(t *testing.T) {
	dictionary := &Dictionary{}
	word := "test"
	err := dictionary.Add(word, "this is just a test definition")
	assert.Nil(t, err)

	err = dictionary.Add(word, "this is a different definition")
	assert.Equal(t, WordExistsError, err)
}

func TestReAddingSameDefinition(t *testing.T) {
	dictionary := &Dictionary{}
	word := "test"
	definition := "this is a dummy definition"

	err := dictionary.Add(word, definition)
	assert.Nil(t, err)
	err = dictionary.Add(word, definition)
	assert.Nil(t, err)
}

func TestUpdateDefinition(t *testing.T) {
	dictionary := &Dictionary{}
	word := "test"
	newDefinition := "new definition"
	err := dictionary.Add(word, "original definition")
	assert.Nil(t, err)
	err = dictionary.Update(word, newDefinition)
	assert.Nil(t, err)

	definition, err := dictionary.Search(word)
	assert.Nil(t, err)
	assert.Equal(t, newDefinition, definition)
}

func TestUpdateWordThatDoesNotExist(t *testing.T) {
	dictionary := &Dictionary{}
	err := dictionary.Update("test", "new definition")
	assert.Equal(t, WordDoesNotExistError, err)
}

func TestDelete(t *testing.T) {
	dictionary := &Dictionary{}
	word := "test"
	err := dictionary.Add(word, "dummy definition")
	assert.Nil(t, err)

	err = dictionary.Delete(word)
	assert.Nil(t, err)
	_, err = dictionary.Search(word)
	assert.Equal(t, NoDefinitionError, err)
}

func TestDeleteWordDoesNotExist(t *testing.T) {
	dictionary := &Dictionary{}
	err := dictionary.Delete("word doesn't exist")
	assert.Equal(t, WordDoesNotExistError, err)
}
