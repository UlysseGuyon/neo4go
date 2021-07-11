package neo4go

import (
	"time"

	internalTypes "github.com/UlysseGuyon/neo4go/internal/types"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func convertInputObject(obj InputObject) interface{} {
	if obj == nil {
		return nil
	}

	if primitive, canConvert := obj.(internalTypes.PrimitiveInputObject); canConvert {
		return primitive.PrimitiveConvert()
	}

	rootMap := obj.Convert()
	interfaceMap := make(map[string]interface{})
	for key, val := range rootMap {
		interfaceMap[key] = convertInputObject(val)
	}

	return interfaceMap
}

type inputArray struct {
	Value []InputObject
}

func NewInputArray(value []InputObject) InputObject {
	return &inputArray{Value: value}
}

func (val *inputArray) Convert() map[string]InputObject {
	return nil
}

func (val *inputArray) PrimitiveConvert() interface{} {
	resArray := make([]interface{}, 0, len(val.Value))

	for _, v := range val.Value {
		resArray = append(resArray, convertInputObject(v))
	}

	return resArray
}

type inputMap struct {
	Value map[string]InputObject
}

func NewInputMap(value map[string]InputObject) InputObject {
	return &inputMap{Value: value}
}

func (val *inputMap) Convert() map[string]InputObject {
	return val.Value
}

type inputInteger struct {
	Value *int64
}

func NewInputInteger(value *int64) InputObject {
	return &inputInteger{Value: value}
}

func (val *inputInteger) Convert() map[string]InputObject {
	return nil
}

func (val *inputInteger) PrimitiveConvert() interface{} {
	return val.Value
}

type inputFloat struct {
	Value *float64
}

func NewInputFloat(value *float64) InputObject {
	return &inputFloat{Value: value}
}

func (val *inputFloat) Convert() map[string]InputObject {
	return nil
}

func (val *inputFloat) PrimitiveConvert() interface{} {
	return val.Value
}

type inputBool struct {
	Value *bool
}

func NewInputBool(value *bool) InputObject {
	return &inputBool{Value: value}
}

func (val *inputBool) Convert() map[string]InputObject {
	return nil
}

func (val *inputBool) PrimitiveConvert() interface{} {
	return val.Value
}

type inputString struct {
	Value *string
}

func NewInputString(value *string) InputObject {
	return &inputString{Value: value}
}

func (val *inputString) Convert() map[string]InputObject {
	return nil
}

func (val *inputString) PrimitiveConvert() interface{} {
	return val.Value
}

type inputByteArray struct {
	Value []byte
}

func NewInputByteArray(value []byte) InputObject {
	return &inputByteArray{Value: value}
}

func (val *inputByteArray) Convert() map[string]InputObject {
	return nil
}

func (val *inputByteArray) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDate struct {
	Value *neo4j.Date
}

func NewInputDate(value *neo4j.Date) InputObject {
	return &inputDate{Value: value}
}

func (val *inputDate) Convert() map[string]InputObject {
	return nil
}

func (val *inputDate) PrimitiveConvert() interface{} {
	return val.Value
}

type inputTime struct {
	Value *neo4j.OffsetTime
}

func NewInputTime(value *neo4j.OffsetTime) InputObject {
	return &inputTime{Value: value}
}

func (val *inputTime) Convert() map[string]InputObject {
	return nil
}

func (val *inputTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputLocalTime struct {
	Value *neo4j.LocalTime
}

func NewInputLocalTime(value *neo4j.LocalTime) InputObject {
	return &inputLocalTime{Value: value}
}

func (val *inputLocalTime) Convert() map[string]InputObject {
	return nil
}

func (val *inputLocalTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDateTime struct {
	Value *time.Time
}

func NewInputDateTime(value *time.Time) InputObject {
	return &inputDateTime{Value: value}
}

func (val *inputDateTime) Convert() map[string]InputObject {
	return nil
}

func (val *inputDateTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputLocalDateTime struct {
	Value *neo4j.LocalDateTime
}

func NewInputLocalDateTime(value *neo4j.LocalDateTime) InputObject {
	return &inputLocalDateTime{Value: value}
}

func (val *inputLocalDateTime) Convert() map[string]InputObject {
	return nil
}

func (val *inputLocalDateTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDuration struct {
	Value *neo4j.Duration
}

func NewInputDuration(value *neo4j.Duration) InputObject {
	return &inputDuration{Value: value}
}

func (val *inputDuration) Convert() map[string]InputObject {
	return nil
}

func (val *inputDuration) PrimitiveConvert() interface{} {
	return val.Value
}

type inputPoint struct {
	Value *neo4j.Point
}

func NewInputPoint(value *neo4j.Point) InputObject {
	return &inputPoint{Value: value}
}

func (val *inputPoint) Convert() map[string]InputObject {
	return nil
}

func (val *inputPoint) PrimitiveConvert() interface{} {
	return val.Value
}

type inputNode struct {
	Value neo4j.Node
}

func NewInputNode(value neo4j.Node) InputObject {
	return &inputNode{Value: value}
}

func (val *inputNode) Convert() map[string]InputObject {
	return nil
}

func (val *inputNode) PrimitiveConvert() interface{} {
	return val.Value
}

type inputRelationship struct {
	Value neo4j.Relationship
}

func NewInputRelationship(value neo4j.Relationship) InputObject {
	return &inputRelationship{Value: value}
}

func (val *inputRelationship) Convert() map[string]InputObject {
	return nil
}

func (val *inputRelationship) PrimitiveConvert() interface{} {
	return val.Value
}

type inputPath struct {
	Value neo4j.Path
}

func NewInputPath(value neo4j.Path) InputObject {
	return &inputPath{Value: value}
}

func (val *inputPath) Convert() map[string]InputObject {
	return nil
}

func (val *inputPath) PrimitiveConvert() interface{} {
	return val.Value
}
