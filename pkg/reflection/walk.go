package reflection

import "reflect"

func getValue(x interface{}) (val reflect.Value) {
	val = reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return
}

func walk(x interface{}, fn func(string)) {
	val := getValue(x)
	numberOfFields := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		numberOfFields = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfFields = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if x, ok := val.Recv(); ok {
				walk(x.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, result := range valFnResult {
			walk(result.Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	default:
	}

	for i := 0; i < numberOfFields; i++ {
		walk(getField(i).Interface(), fn)
	}
}
