package walk

import "reflect"

func walk(input interface{}, action func(string)) {
	val := reflect.ValueOf(input)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			action(field.String())
		}
	}
}
