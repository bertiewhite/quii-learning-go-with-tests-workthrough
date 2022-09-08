package greetings

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	greeting := fmt.Sprintf("Hello, %v", name)
	writer.Write([]byte(greeting))
}
