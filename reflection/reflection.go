package reflection_

import "reflect"

func Walk(x any, fn func(input string)) {
	val := getValue(x)
	valLength := 0

	var getField func(i int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		valLength = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		valLength = val.Len()
		getField = val.Index
	}

	for i := 0; i < valLength; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
