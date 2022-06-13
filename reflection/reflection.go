package reflection_

import "reflect"

func Walk(x any, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
