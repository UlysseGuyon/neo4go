package neo4go

import "reflect"

func GetValueElem(val reflect.Value) reflect.Value {
	resVal := val

	if !resVal.IsValid() {
		return resVal
	}

	if resVal.Kind() == reflect.Ptr || resVal.Kind() == reflect.Interface {
		if resVal.IsNil() {
			return resVal
		}
		resVal = resVal.Elem()
	}

	return resVal
}
