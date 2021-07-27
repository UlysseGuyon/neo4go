package neo4go

import "reflect"

// GetValueElem returns the underlying value of a reflected value and passes through pointers/interfaces
func GetValueElem(val reflect.Value) reflect.Value {
	resVal := val

	for resVal.Kind() == reflect.Ptr || resVal.Kind() == reflect.Interface {
		if !resVal.IsValid() {
			return resVal
		}

		if resVal.IsNil() {
			return resVal
		}

		resVal = resVal.Elem()
	}

	return resVal
}
