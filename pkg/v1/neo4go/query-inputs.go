package neo4go

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// InputStruct represents a struct as an acceptable input for the manager queries.
// It is also the most basic level of query input
type InputStruct interface {
	// ConvertToMap converts this input as a map of query inputs
	ConvertToMap() map[string]InputStruct
}

// InputOtherType represents a any complex type as an acceptable input for the manager queries.
type InputOtherType interface {
	InputStruct

	// ConvertToInputObject directly converts the object as an input object?
	ConvertToInputObject() InputStruct
}

// primitiveInputObject represents a go primitive value converted as an input object for the manager queries.
type primitiveInputObject interface {
	InputStruct
	InputOtherType

	// PrimitiveConvert directly converts the object as an interface and
	// should not be used outside of this package to allow fully functionning type checking
	PrimitiveConvert() interface{}
}

// convertInputObject takes an input object and converts it into an interface using firstly the most primitive object functions
func convertInputObject(obj InputStruct) interface{} {
	if obj == nil {
		return nil
	}

	// First, check for primitive value. This is the stop condition of this recursive function
	if primitive, canConvert := obj.(primitiveInputObject); canConvert {
		return primitive.PrimitiveConvert()
	}

	// Then, check for non struct value
	if convertedPrimitive, canConvert := obj.(InputOtherType); canConvert {
		inObj := convertedPrimitive.ConvertToInputObject()
		return convertInputObject(inObj)
	}

	// Else, it is a struct input
	rootMap := obj.ConvertToMap()
	interfaceMap := make(map[string]interface{})
	for key, val := range rootMap {
		interfaceMap[key] = convertInputObject(val)
	}

	return interfaceMap
}

// inputArray is an implementation of the primitiveInputObject for the neo4j Array type
type inputArray struct {
	Value []InputStruct
}

// NewInputArray creates a primitiveInputObject from the golang type Array
func NewInputArray(value []InputStruct) InputStruct {
	return &inputArray{Value: value}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputArray) ConvertToMap() map[string]InputStruct {
	return nil
}

// ConvertToInputObject directly converts the object as an input object?
func (val *inputArray) ConvertToInputObject() InputStruct {
	return val
}

// PrimitiveConvert directly converts the object as an interface and
// should not be used outside of this package to allow fully functionning type checking
func (val *inputArray) PrimitiveConvert() interface{} {
	resArray := make([]interface{}, 0, len(val.Value))

	for _, v := range val.Value {
		resArray = append(resArray, convertInputObject(v))
	}

	return resArray
}

// inputMap is an implementation of the InputStruct for the neo4j Map type
type inputMap struct {
	Value map[string]InputStruct
}

// NewInputMap creates a InputStruct from the golang type Map
func NewInputMap(value map[string]InputStruct) InputStruct {
	return &inputMap{Value: value}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputMap) ConvertToMap() map[string]InputStruct {
	return val.Value
}

// inputInteger is an implementation of the primitiveInputObject for the neo4j Integer type
type inputInteger struct {
	Value *int64
}

// NewInputInteger creates a primitiveInputObject from the golang type Integer
func NewInputInteger(value *int64) InputStruct {
	return &inputInteger{Value: value}
}

// NewInputUnsignedInteger creates a primitiveInputObject from the golang type Unsigned Integer
func NewInputUnsignedInteger(value *uint64) InputStruct {
	var convertedInt *int64
	if value != nil {
		convertedRaw := int64(*value)
		convertedInt = &(convertedRaw)
	}
	return &inputInteger{Value: convertedInt}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputInteger) ConvertToMap() map[string]InputStruct {
	return nil
}

// ConvertToInputObject directly converts the object as an input object?
func (val *inputInteger) ConvertToInputObject() InputStruct {
	return val
}

// PrimitiveConvert directly converts the object as an interface and
// should not be used outside of this package to allow fully functionning type checking
func (val *inputInteger) PrimitiveConvert() interface{} {
	return val.Value
}

// inputFloat is an implementation of the primitiveInputObject for the neo4j Float type
type inputFloat struct {
	Value *float64
}

// NewInputFloat creates a primitiveInputObject from the golang type Float
func NewInputFloat(value *float64) InputStruct {
	return &inputFloat{Value: value}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputFloat) ConvertToMap() map[string]InputStruct {
	return nil
}

// ConvertToInputObject directly converts the object as an input object?
func (val *inputFloat) ConvertToInputObject() InputStruct {
	return val
}

// PrimitiveConvert directly converts the object as an interface and
// should not be used outside of this package to allow fully functionning type checking
func (val *inputFloat) PrimitiveConvert() interface{} {
	return val.Value
}

// inputBool is an implementation of the primitiveInputObject for the neo4j Bool type
type inputBool struct {
	Value *bool
}

// NewInputBool creates a primitiveInputObject from the golang type Bool
func NewInputBool(value *bool) InputStruct {
	return &inputBool{Value: value}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputBool) ConvertToMap() map[string]InputStruct {
	return nil
}

// ConvertToInputObject directly converts the object as an input object?
func (val *inputBool) ConvertToInputObject() InputStruct {
	return val
}

// PrimitiveConvert directly converts the object as an interface and
// should not be used outside of this package to allow fully functionning type checking
func (val *inputBool) PrimitiveConvert() interface{} {
	return val.Value
}

// inputString is an implementation of the primitiveInputObject for the neo4j String type
type inputString struct {
	Value *string
}

// NewInputString creates a primitiveInputObject from the golang type String
func NewInputString(value *string) InputStruct {
	return &inputString{Value: value}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputString) ConvertToMap() map[string]InputStruct {
	return nil
}

// ConvertToInputObject directly converts the object as an input object?
func (val *inputString) ConvertToInputObject() InputStruct {
	return val
}

// PrimitiveConvert directly converts the object as an interface and
// should not be used outside of this package to allow fully functionning type checking
func (val *inputString) PrimitiveConvert() interface{} {
	return val.Value
}

// inputByteArray is an implementation of the primitiveInputObject for the neo4j ByteArray type
type inputByteArray struct {
	Value []byte
}

// NewInputByteArray creates a primitiveInputObject from the golang type []byte
func NewInputByteArray(value []byte) InputStruct {
	return &inputByteArray{Value: value}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputByteArray) ConvertToMap() map[string]InputStruct {
	return nil
}

// ConvertToInputObject directly converts the object as an input object?
func (val *inputByteArray) ConvertToInputObject() InputStruct {
	return val
}

// PrimitiveConvert directly converts the object as an interface and
// should not be used outside of this package to allow fully functionning type checking
func (val *inputByteArray) PrimitiveConvert() interface{} {
	return val.Value
}

// inputDateTime is an implementation of the primitiveInputObject for the neo4j DateTime type
type inputDateTime struct {
	Value *time.Time
}

// NewInputDateTime creates a primitiveInputObject from the golang type Time
func NewInputDateTime(value *time.Time) InputStruct {
	return &inputDateTime{Value: value}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputDateTime) ConvertToMap() map[string]InputStruct {
	return nil
}

// ConvertToInputObject directly converts the object as an input object?
func (val *inputDateTime) ConvertToInputObject() InputStruct {
	return val
}

// PrimitiveConvert directly converts the object as an interface and
// should not be used outside of this package to allow fully functionning type checking
func (val *inputDateTime) PrimitiveConvert() interface{} {
	if val.Value == nil {
		return nil
	} else {
		return *(val.Value)
	}
}

// inputPoint is an implementation of the primitiveInputObject for the neo4j Point type
type inputPoint struct {
	Value *neo4j.Point
}

// NewInputPoint creates a primitiveInputObject from the golang type neo4j.Point
func NewInputPoint(value *neo4j.Point) InputStruct {
	return &inputPoint{Value: value}
}

// ConvertToMap converts this input as a map of query inputs
func (val *inputPoint) ConvertToMap() map[string]InputStruct {
	x := val.Value.X()
	y := val.Value.Y()
	z := val.Value.Z()
	srid := int64(val.Value.SrId())

	return map[string]InputStruct{
		"x":    NewInputFloat(&x),
		"y":    NewInputFloat(&y),
		"z":    NewInputFloat(&z),
		"srid": NewInputInteger(&srid),
	}
}
