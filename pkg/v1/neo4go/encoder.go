package neo4go

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	internalMain "github.com/UlysseGuyon/neo4go/internal/neo4go"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Neo4GoEncoder interface {
	Encode(interface{}) InputStruct
}

type EncodeHookFunc func(reflect.Value, interface{}) (InputStruct, bool)

type Neo4GoEncoderOptions struct {
	TagName    string
	EncodeHook EncodeHookFunc
}

type neo4goEncoder struct {
	Options Neo4GoEncoderOptions
}

func NewNeo4GoEncoder(opt *Neo4GoEncoderOptions) Neo4GoEncoder {
	usedOpt := Neo4GoEncoderOptions{}
	if opt != nil {
		usedOpt = *opt
	}

	newEncoder := neo4goEncoder{
		Options: usedOpt,
	}

	if newEncoder.Options.TagName == "" {
		newEncoder.Options.TagName = internalMain.DefaultEncodingTagName
	}

	var usedHook EncodeHookFunc
	if newEncoder.Options.EncodeHook == nil {
		usedHook = newEncoder.getDefaultHook()
	} else {
		usedHook = ComposeEncodeHookFunc(
			defaultHookNil, // NOTE This one must be first in order to detect nil values without panic
			newEncoder.Options.EncodeHook,
			newEncoder.getDefaultHook(),
		)
	}

	newEncoder.Options.EncodeHook = usedHook

	return &newEncoder
}

func (encoder *neo4goEncoder) Encode(obj interface{}) InputStruct {
	objValue := reflect.ValueOf(obj)

	if encodedObj, canEncode := encoder.Options.EncodeHook(objValue, obj); canEncode {
		return encodedObj
	}

	return nil
}

func (encoder *neo4goEncoder) getDefaultHook() EncodeHookFunc {
	return ComposeEncodeHookFunc(
		defaultHookInputStruct, // NOTE This one must be first in order to have the wanted behavior
		defaultHookInteger,
		defaultHookFloat,
		defaultHookBool,
		defaultHookString,
		defaultHookByteArray,
		defaultHookTime,
		defaultHookDuration,
		defaultHookPoint,
		defaultHookNode,
		defaultHookRelationship,
		defaultHookPath,
		defaultHookStruct(encoder.Options.TagName, encoder),
		defaultHookArray(encoder),
		defaultHookMap(encoder),
	)
}

func ComposeEncodeHookFunc(hooks ...EncodeHookFunc) EncodeHookFunc {
	return func(v reflect.Value, i interface{}) (InputStruct, bool) {
		for _, hook := range hooks {
			inputStruct, converted := hook(v, i)
			if converted {
				return inputStruct, converted
			}
		}

		return nil, false
	}
}

// All default hooks
var (
	defaultHookNil EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if i == nil || !v.IsValid() {
			return nil, true
		}

		return nil, false
	}

	defaultHookInputStruct EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if inputStruct, canConvert := i.(InputStruct); canConvert {
			return inputStruct, true
		}

		return nil, false
	}

	defaultHookInteger EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		usedVal := GetValueElem(v)

		switch usedVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			objInt := usedVal.Int()
			return NewInputInteger(&objInt), true
		default:
			return nil, false
		}
	}

	defaultHookFloat EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		usedVal := GetValueElem(v)

		switch usedVal.Kind() {
		case reflect.Float32, reflect.Float64:
			objInt := usedVal.Float()
			return NewInputFloat(&objInt), true
		default:
			return nil, false
		}
	}

	defaultHookBool EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		usedVal := GetValueElem(v)

		switch usedVal.Kind() {
		case reflect.Bool:
			objInt := usedVal.Bool()
			return NewInputBool(&objInt), true
		default:
			return nil, false
		}
	}

	defaultHookString EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		usedVal := GetValueElem(v)

		switch usedVal.Kind() {
		case reflect.String:
			objInt := usedVal.String()
			return NewInputString(&objInt), true
		default:
			return nil, false
		}
	}

	defaultHookByteArray EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if byteArr, canConvert := i.([]byte); canConvert {
			return NewInputByteArray(byteArr), true
		}

		return nil, false
	}

	defaultHookTime EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if timeVal, canConvert := i.(time.Time); canConvert {
			timeValNeo4j := neo4j.LocalDateTimeOf(timeVal)
			return NewInputLocalDateTime(&timeValNeo4j), true
		}

		return nil, false
	}

	defaultHookDuration EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if duration, canConvert := i.(time.Duration); canConvert {
			durationValNeo4j := neo4j.DurationOf(0, int64(duration.Hours()/24), int64(duration.Seconds()), int(duration.Nanoseconds()))
			return NewInputDuration(&durationValNeo4j), true
		}

		return nil, false
	}

	defaultHookPoint EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		var usedPoint neo4j.Point

		switch typedI := i.(type) {
		case *neo4j.Point:
			if typedI != nil {
				usedPoint = *typedI
			} else {
				return nil, false
			}
		case neo4j.Point:
			usedPoint = typedI
		default:
			return nil, false
		}

		return NewInputPoint(&usedPoint), true
	}

	defaultHookNode EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		var usedNode neo4j.Node

		switch typedI := i.(type) {
		case *neo4j.Node:
			if typedI != nil {
				usedNode = *typedI
			} else {
				return nil, false
			}
		case neo4j.Node:
			usedNode = typedI
		default:
			return nil, false
		}

		return NewInputNode(usedNode), true
	}

	defaultHookRelationship EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		var usedRelationship neo4j.Relationship

		switch typedI := i.(type) {
		case *neo4j.Relationship:
			if typedI != nil {
				usedRelationship = *typedI
			} else {
				return nil, false
			}
		case neo4j.Relationship:
			usedRelationship = typedI
		default:
			return nil, false
		}

		return NewInputRelationship(usedRelationship), true
	}

	defaultHookPath EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		var usedPath neo4j.Path

		switch typedI := i.(type) {
		case *neo4j.Path:
			if typedI != nil {
				usedPath = *typedI
			} else {
				return nil, false
			}
		case neo4j.Path:
			usedPath = typedI
		default:
			return nil, false
		}

		return NewInputPath(usedPath), true
	}

	defaultHookStruct = func(tagName string, encoder Neo4GoEncoder) EncodeHookFunc {
		return func(v reflect.Value, i interface{}) (InputStruct, bool) {
			usedVal := GetValueElem(v)
			if usedVal.Kind() != reflect.Struct {
				return nil, false
			}

			usedType := usedVal.Type()
			resultMap := make(map[string]InputStruct)

			for i := 0; i < usedVal.NumField(); i++ {
				field := usedType.Field(i)
				key := strings.ToLower(field.Name)

				fieldVal := usedVal.FieldByName(field.Name)
				fieldTag := field.Tag.Get(tagName)

				if fieldTag == "" {
					continue
				}

				allTagValues := strings.Split(fieldTag, ",")
				nameInTag := strings.TrimSpace(allTagValues[0])
				hasOmitEmpty := false
				for _, tagValue := range allTagValues {
					if tagValue == "omitempty" {
						hasOmitEmpty = true
					}
				}

				if nameInTag != "" && nameInTag != "-" {
					key = nameInTag
				}

				// encode here if the field is exported
				fieldVal = GetValueElem(fieldVal)
				if fieldVal.IsValid() {
					if fieldVal.IsZero() && hasOmitEmpty {
						continue
					} else if fieldVal.CanInterface() {
						fieldInterface := fieldVal.Interface()

						resultMap[key] = encoder.Encode(fieldInterface)
					} else {
						resultMap[key] = nil
					}
				}
			}

			return NewInputMap(resultMap), true
		}
	}

	defaultHookArray = func(encoder Neo4GoEncoder) EncodeHookFunc {
		return func(v reflect.Value, i interface{}) (InputStruct, bool) {
			usedVal := GetValueElem(v)

			if usedVal.Kind() != reflect.Array && usedVal.Kind() != reflect.Slice {
				return nil, false
			}

			encodedArray := make([]InputStruct, 0, usedVal.Len())
			for index := 0; index < usedVal.Len(); index++ {
				itemVal := usedVal.Index(index)
				if itemVal.CanInterface() {
					encodedArray = append(encodedArray, encoder.Encode(itemVal.Interface()))
				} else {
					encodedArray = append(encodedArray, nil)
				}
			}

			return NewInputArray(encodedArray), true
		}
	}

	defaultHookMap = func(encoder Neo4GoEncoder) EncodeHookFunc {
		return func(v reflect.Value, i interface{}) (InputStruct, bool) {
			usedVal := GetValueElem(v)

			if usedVal.Kind() != reflect.Map {
				return nil, false
			}

			encodedMap := make(map[string]InputStruct)
			vMapIter := usedVal.MapRange()
			for vMapIter.Next() {
				key := vMapIter.Key()
				val := vMapIter.Value()

				var keyStr string
				switch key.Type().Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					keyStr = strconv.Itoa(int(key.Int()))
				case reflect.Float32, reflect.Float64:
					keyStr = strconv.FormatFloat(key.Float(), 'f', -1, 64)
				case reflect.String:
					keyStr = key.String()
				default:
					keyStr = key.String()
				}

				if val.CanInterface() {
					encodedMap[keyStr] = encoder.Encode(val.Interface())
				} else {
					encodedMap[keyStr] = nil
				}
			}

			return NewInputMap(encodedMap), true
		}
	}
)
