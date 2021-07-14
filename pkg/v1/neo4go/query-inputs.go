package neo4go

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type InputStruct interface {
	ConvertToMap() map[string]InputStruct
}

type InputOtherType interface {
	InputStruct
	ConvertToInputObject() InputStruct
}

type primitiveInputObject interface {
	InputStruct
	InputOtherType
	PrimitiveConvert() interface{}
}

func convertInputObject(obj InputStruct) interface{} {
	if obj == nil {
		return nil
	}

	if primitive, canConvert := obj.(primitiveInputObject); canConvert {
		return primitive.PrimitiveConvert()
	}

	if convertedPrimitive, canConvert := obj.(InputOtherType); canConvert {
		inObj := convertedPrimitive.ConvertToInputObject()
		return convertInputObject(inObj)
	}

	rootMap := obj.ConvertToMap()
	interfaceMap := make(map[string]interface{})
	for key, val := range rootMap {
		interfaceMap[key] = convertInputObject(val)
	}

	return interfaceMap
}

type inputArray struct {
	Value []InputStruct
}

func NewInputArray(value []InputStruct) primitiveInputObject {
	return &inputArray{Value: value}
}

func (val *inputArray) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputArray) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputArray) PrimitiveConvert() interface{} {
	resArray := make([]interface{}, 0, len(val.Value))

	for _, v := range val.Value {
		resArray = append(resArray, convertInputObject(v))
	}

	return resArray
}

type inputMap struct {
	Value map[string]InputStruct
}

func NewInputMap(value map[string]InputStruct) InputStruct {
	return &inputMap{Value: value}
}

func (val *inputMap) ConvertToMap() map[string]InputStruct {
	return val.Value
}

type inputInteger struct {
	Value *int64
}

func NewInputInteger(value *int64) primitiveInputObject {
	return &inputInteger{Value: value}
}

func (val *inputInteger) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputInteger) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputInteger) PrimitiveConvert() interface{} {
	return val.Value
}

type inputFloat struct {
	Value *float64
}

func NewInputFloat(value *float64) primitiveInputObject {
	return &inputFloat{Value: value}
}

func (val *inputFloat) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputFloat) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputFloat) PrimitiveConvert() interface{} {
	return val.Value
}

type inputBool struct {
	Value *bool
}

func NewInputBool(value *bool) primitiveInputObject {
	return &inputBool{Value: value}
}

func (val *inputBool) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputBool) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputBool) PrimitiveConvert() interface{} {
	return val.Value
}

type inputString struct {
	Value *string
}

func NewInputString(value *string) primitiveInputObject {
	return &inputString{Value: value}
}

func (val *inputString) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputString) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputString) PrimitiveConvert() interface{} {
	return val.Value
}

type inputByteArray struct {
	Value []byte
}

func NewInputByteArray(value []byte) primitiveInputObject {
	return &inputByteArray{Value: value}
}

func (val *inputByteArray) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputByteArray) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputByteArray) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDate struct {
	Value *neo4j.Date
}

func NewInputDate(value *neo4j.Date) primitiveInputObject {
	return &inputDate{Value: value}
}

func (val *inputDate) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputDate) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputDate) PrimitiveConvert() interface{} {
	return val.Value
}

type inputTime struct {
	Value *neo4j.OffsetTime
}

func NewInputTime(value *neo4j.OffsetTime) primitiveInputObject {
	return &inputTime{Value: value}
}

func (val *inputTime) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputTime) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputLocalTime struct {
	Value *neo4j.LocalTime
}

func NewInputLocalTime(value *neo4j.LocalTime) primitiveInputObject {
	return &inputLocalTime{Value: value}
}

func (val *inputLocalTime) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputLocalTime) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputLocalTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDateTime struct {
	Value *time.Time
}

func NewInputDateTime(value *time.Time) primitiveInputObject {
	return &inputDateTime{Value: value}
}

func (val *inputDateTime) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputDateTime) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputDateTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputLocalDateTime struct {
	Value *neo4j.LocalDateTime
}

func (val *inputLocalDateTime) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputLocalDateTime) ConvertToInputObject() InputStruct {
	return val
}

func NewInputLocalDateTime(value *neo4j.LocalDateTime) primitiveInputObject {
	return &inputLocalDateTime{Value: value}
}

func (val *inputLocalDateTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDuration struct {
	Value *neo4j.Duration
}

func (val *inputDuration) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputDuration) ConvertToInputObject() InputStruct {
	return val
}

func NewInputDuration(value *neo4j.Duration) primitiveInputObject {
	return &inputDuration{Value: value}
}

func (val *inputDuration) PrimitiveConvert() interface{} {
	return val.Value
}

type inputPoint struct {
	Value *neo4j.Point
}

func NewInputPoint(value *neo4j.Point) primitiveInputObject {
	return &inputPoint{Value: value}
}

func (val *inputPoint) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputPoint) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputPoint) PrimitiveConvert() interface{} {
	return val.Value
}

type inputNode struct {
	Value neo4j.Node
}

func NewInputNode(value neo4j.Node) primitiveInputObject {
	return &inputNode{Value: value}
}

func (val *inputNode) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputNode) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputNode) PrimitiveConvert() interface{} {
	return val.Value
}

type inputRelationship struct {
	Value neo4j.Relationship
}

func NewInputRelationship(value neo4j.Relationship) primitiveInputObject {
	return &inputRelationship{Value: value}
}

func (val *inputRelationship) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputRelationship) ConvertToInputObject() InputStruct {
	return val
}

func (val *inputRelationship) PrimitiveConvert() interface{} {
	return val.Value
}

type inputPath struct {
	Value neo4j.Path
}

func (val *inputPath) ConvertToMap() map[string]InputStruct {
	return nil
}

func (val *inputPath) ConvertToInputObject() InputStruct {
	return val
}

func NewInputPath(value neo4j.Path) primitiveInputObject {
	return &inputPath{Value: value}
}

func (val *inputPath) PrimitiveConvert() interface{} {
	return val.Value
}
