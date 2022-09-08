package greetings

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}
	Greet(buffer, "Chris")

	expGreeting := "Hello, Chris"
	assert.Equal(t, expGreeting, buffer.String())
}
