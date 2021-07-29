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

// IsNil tells if the value can be used or not (as a plain object or as a dereferenceable pointer)
func IsNil(val reflect.Value) bool {
	if !val.IsValid() {
		return true
	}

	if val.Kind() == reflect.Chan || val.Kind() == reflect.Func || val.Kind() == reflect.Interface ||
		val.Kind() == reflect.Map || val.Kind() == reflect.Ptr || val.Kind() == reflect.Slice {
		return val.IsNil()
	}

	return false
}
