package neo4go

import (
	"time"

	internalTypes "github.com/UlysseGuyon/neo4go/internal/types"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func convertInputObject(obj interface{}) interface{} {
	if obj == nil {
		return nil
	}

	if structObj, canConvert := obj.(InputStruct); canConvert {
		rootMap := structObj.ConvertToMap()
		interfaceMap := make(map[string]interface{})
		for key, val := range rootMap {
			interfaceMap[key] = convertInputObject(val)
		}
	}

	if convertedPrimitive, canConvert := obj.(InputPrimitive); canConvert {
		return convertedPrimitive.ConvertToPrimitive().PrimitiveConvert()
	}

	if primitive, canConvert := obj.(internalTypes.PrimitiveInputObject); canConvert {
		return primitive.PrimitiveConvert()
	}

	return nil
}

type inputArray struct {
	Value []InputStruct
}

func NewInputArray(value []InputStruct) internalTypes.PrimitiveInputObject {
	return &inputArray{Value: value}
}

func (val *inputArray) PrimitiveConvert() interface{} {
	resArray := make([]interface{}, 0, len(val.Value))

	for _, v := range val.Value {
		resArray = append(resArray, convertInputObject(v))
	}

	return resArray
}

type inputMap struct {
	Value map[string]internalTypes.InputStruct
}

func NewInputMap(value map[string]InputStruct) internalTypes.InputStruct {
	convertedValMap := make(map[string]internalTypes.InputStruct, len(value))

	for key, val := range value {
		convertedVal := val.(internalTypes.InputStruct)
		convertedValMap[key] = convertedVal
	}

	return &inputMap{Value: convertedValMap}
}

func (val *inputMap) ConvertToMap() map[string]internalTypes.InputStruct {
	return val.Value
}

type inputInteger struct {
	Value *int64
}

func NewInputInteger(value *int64) internalTypes.PrimitiveInputObject {
	return &inputInteger{Value: value}
}

func (val *inputInteger) PrimitiveConvert() interface{} {
	return val.Value
}

type inputFloat struct {
	Value *float64
}

func NewInputFloat(value *float64) internalTypes.PrimitiveInputObject {
	return &inputFloat{Value: value}
}

func (val *inputFloat) PrimitiveConvert() interface{} {
	return val.Value
}

type inputBool struct {
	Value *bool
}

func NewInputBool(value *bool) internalTypes.PrimitiveInputObject {
	return &inputBool{Value: value}
}

func (val *inputBool) PrimitiveConvert() interface{} {
	return val.Value
}

type inputString struct {
	Value *string
}

func NewInputString(value *string) internalTypes.PrimitiveInputObject {
	return &inputString{Value: value}
}

func (val *inputString) PrimitiveConvert() interface{} {
	return val.Value
}

type inputByteArray struct {
	Value []byte
}

func NewInputByteArray(value []byte) internalTypes.PrimitiveInputObject {
	return &inputByteArray{Value: value}
}

func (val *inputByteArray) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDate struct {
	Value *neo4j.Date
}

func NewInputDate(value *neo4j.Date) internalTypes.PrimitiveInputObject {
	return &inputDate{Value: value}
}

func (val *inputDate) PrimitiveConvert() interface{} {
	return val.Value
}

type inputTime struct {
	Value *neo4j.OffsetTime
}

func NewInputTime(value *neo4j.OffsetTime) internalTypes.PrimitiveInputObject {
	return &inputTime{Value: value}
}

func (val *inputTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputLocalTime struct {
	Value *neo4j.LocalTime
}

func NewInputLocalTime(value *neo4j.LocalTime) internalTypes.PrimitiveInputObject {
	return &inputLocalTime{Value: value}
}

func (val *inputLocalTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDateTime struct {
	Value *time.Time
}

func NewInputDateTime(value *time.Time) internalTypes.PrimitiveInputObject {
	return &inputDateTime{Value: value}
}

func (val *inputDateTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputLocalDateTime struct {
	Value *neo4j.LocalDateTime
}

func NewInputLocalDateTime(value *neo4j.LocalDateTime) internalTypes.PrimitiveInputObject {
	return &inputLocalDateTime{Value: value}
}

func (val *inputLocalDateTime) PrimitiveConvert() interface{} {
	return val.Value
}

type inputDuration struct {
	Value *neo4j.Duration
}

func NewInputDuration(value *neo4j.Duration) internalTypes.PrimitiveInputObject {
	return &inputDuration{Value: value}
}

func (val *inputDuration) PrimitiveConvert() interface{} {
	return val.Value
}

type inputPoint struct {
	Value *neo4j.Point
}

func NewInputPoint(value *neo4j.Point) internalTypes.PrimitiveInputObject {
	return &inputPoint{Value: value}
}

func (val *inputPoint) PrimitiveConvert() interface{} {
	return val.Value
}

type inputNode struct {
	Value neo4j.Node
}

func NewInputNode(value neo4j.Node) internalTypes.PrimitiveInputObject {
	return &inputNode{Value: value}
}

func (val *inputNode) PrimitiveConvert() interface{} {
	return val.Value
}

type inputRelationship struct {
	Value neo4j.Relationship
}

func NewInputRelationship(value neo4j.Relationship) internalTypes.PrimitiveInputObject {
	return &inputRelationship{Value: value}
}

func (val *inputRelationship) PrimitiveConvert() interface{} {
	return val.Value
}

type inputPath struct {
	Value neo4j.Path
}

func NewInputPath(value neo4j.Path) internalTypes.PrimitiveInputObject {
	return &inputPath{Value: value}
}

func (val *inputPath) PrimitiveConvert() interface{} {
	return val.Value
}
