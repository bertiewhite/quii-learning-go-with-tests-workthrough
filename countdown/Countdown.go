package countdown

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer, sleeper Sleeper) {
	max := 3
	for i := max; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(writer, "Go!")
}
