package walk

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"chris"},
			[]string{"chris"},
		},
		{
			"Struct with 2 strings",
			struct {
				Name string
				city string
			}{"Stephen", "Bristol"},
			[]string{"Stephen", "Bristol"},
		},
		{
			"Struct with one non-string field",
			struct {
				Name string
				Age  int
			}{"Jermaine", 3},
			[]string{"Jermaine"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var received []string

			walk(test.Input, func(s string) {
				received = append(received, s)
			})

			assert.True(t, reflect.DeepEqual(received, test.ExpectedCalls))
		})
	}
}
