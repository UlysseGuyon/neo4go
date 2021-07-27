package neo4go

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	internalMain "github.com/UlysseGuyon/neo4go/internal/neo4go"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Encoder allows a user to encode any type of data (including custom structs) into neo4go query input values
type Encoder interface {
	// Encode takes any object and encodes it into an object accepted by the neo4go query system
	Encode(interface{}) InputStruct
}

// EncodeHookFunc represents a function that converts a specific type of value into a neo4go query input
type EncodeHookFunc func(reflect.Value, interface{}) (InputStruct, bool)

// EncoderOptions represents the configuration applied to an encoder
type EncoderOptions struct {
	// The tag name used to find and encode struct fields
	TagName string

	// The function used for every input of this encoder. It is typically a composition of other functions
	EncodeHook EncodeHookFunc

	// Tells if the encoder should be silent or not when it finds an object it cannot decode
	Silent bool
}

// neo4goEncoder is the default implementation of the Encoder interface
type neo4goEncoder struct {
	// The configuration of this encoder
	options EncoderOptions
}

// NewEncoder creates a new instance of Encoder, with a given config. A nil config will result in the default config beinng applied
func NewEncoder(opt *EncoderOptions) Encoder {
	// Use the given config if not nil
	usedOpt := EncoderOptions{}
	if opt != nil {
		usedOpt = *opt
	}

	newEncoder := neo4goEncoder{
		options: usedOpt,
	}

	// Use the default encoding tag name if none is given
	if newEncoder.options.TagName == "" {
		newEncoder.options.TagName = internalMain.DefaultEncodingTagName
	}

	// Set the encoding hook as first the nil detector, then the custom hooks, then the default hooks
	// As the ComposeEncodeHookFunc begins the calls by the beggining of the hook list and stops at the first that succeeds
	newEncoder.options.EncodeHook = ComposeEncodeHookFunc(
		defaultHookNil,                // NOTE This one must be first in order to detect nil values without panic
		newEncoder.options.EncodeHook, // NOTE This one must be before the default hooks so that it won't be overriden
		newEncoder.getDefaultHook(),
	)

	return &newEncoder
}

// Encode takes any object and encodes it into an object accepted by the neo4go query system
func (encoder *neo4goEncoder) Encode(obj interface{}) InputStruct {
	objValue := reflect.ValueOf(obj)

	// Call the hook of the encoder and log if it could not encode the object
	if encodedObj, canEncode := encoder.options.EncodeHook(objValue, obj); canEncode {
		return encodedObj
	}

	if !encoder.options.Silent {
		log.Printf("Could not encode object : (Type : %s) %+v\n", objValue.Type().String(), obj)
	}

	return nil
}

// getDefaultHook returns a composition of all the default encode hook functions used for primitive values and array/map/struct
func (encoder *neo4goEncoder) getDefaultHook() EncodeHookFunc {
	return ComposeEncodeHookFunc(
		defaultHookInputStruct, // NOTE This one must be first in order to have the wanted behavior
		defaultHookInteger,
		defaultHookFloat,
		defaultHookBool,
		defaultHookString,
		defaultHookByteArray,
		defaultHookDateTime,
		defaultHookDate,
		defaultHookTime,
		defaultHookLocalTime,
		defaultHookLocalDateTime,
		defaultHookDuration,
		defaultHookGoDuration,
		defaultHookPoint,
		defaultHookNode,
		defaultHookRelationship,
		defaultHookPath,
		defaultHookStruct(encoder.options.TagName, encoder),
		defaultHookArray(encoder),
		defaultHookMap(encoder),
	)
}

// ComposeEncodeHookFunc allows to compose multiple encoding functions into one in order to pass it to an Encoder
func ComposeEncodeHookFunc(hooks ...EncodeHookFunc) EncodeHookFunc {
	return func(v reflect.Value, i interface{}) (InputStruct, bool) {
		for _, hook := range hooks {
			if hook == nil {
				continue
			}

			// For each hook func, call it and end the loop if it could successfully encode the object
			inputStruct, converted := hook(v, i)
			if converted {
				return inputStruct, converted
			}
		}

		return nil, false
	}
}

// All default hooks for primitive and classic encoding
var (
	// The hook that detects potentially problematic values
	defaultHookNil EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if i == nil || !v.IsValid() {
			return nil, true
		}

		return nil, false
	}

	// The hook that does nothing if the object is already a query input value
	defaultHookInputStruct EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if inputStruct, canConvert := i.(InputStruct); canConvert {
			return inputStruct, true
		}

		return nil, false
	}

	// The hook that encodes integer primitive values
	defaultHookInteger EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		usedVal := GetValueElem(v)

		switch usedVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			objInt := usedVal.Int()
			return NewInputInteger(&objInt), true
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			objInt := usedVal.Uint()
			return NewInputUnsignedInteger(&objInt), true
		default:
			return nil, false
		}
	}

	// The hook that encodes float primitive values
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

	// The hook that encodes boolean primitive values
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

	// The hook that encodes string primitive values
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

	// The hook that encodes byte array primitive values
	defaultHookByteArray EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if byteArr, canConvert := i.([]byte); canConvert {
			return NewInputByteArray(byteArr), true
		}

		return nil, false
	}

	// The hook that encodes time values as local datetimes. This will always store time values as UTC
	defaultHookDateTime EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if timeVal, canConvert := i.(time.Time); canConvert {
			return NewInputUTCTime(&timeVal), true
		} else if timeVal, canConvert := i.(*time.Time); canConvert {
			return NewInputUTCTime(timeVal), true
		}

		return nil, false
	}

	// The hook that encodes neo4j dates
	defaultHookDate EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if dateVal, canConvert := i.(neo4j.Date); canConvert {
			return NewInputDate(&dateVal), true
		} else if dateVal, canConvert := i.(*neo4j.Date); canConvert {
			return NewInputDate(dateVal), true
		}

		return nil, false
	}

	// The hook that encodes neo4j offset times
	defaultHookTime EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if oofsetTimeVal, canConvert := i.(neo4j.OffsetTime); canConvert {
			return NewInputTime(&oofsetTimeVal), true
		} else if oofsetTimeVal, canConvert := i.(*neo4j.OffsetTime); canConvert {
			return NewInputTime(oofsetTimeVal), true
		}

		return nil, false
	}

	// The hook that encodes neo4j local times
	defaultHookLocalTime EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if localTimeVal, canConvert := i.(neo4j.LocalTime); canConvert {
			return NewInputLocalTime(&localTimeVal), true
		} else if localTimeVal, canConvert := i.(*neo4j.LocalTime); canConvert {
			return NewInputLocalTime(localTimeVal), true
		}

		return nil, false
	}

	// The hook that encodes neo4j local times
	defaultHookLocalDateTime EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if localDateTimeVal, canConvert := i.(neo4j.LocalDateTime); canConvert {
			return NewInputLocalDateTime(&localDateTimeVal), true
		} else if localDateTimeVal, canConvert := i.(*neo4j.LocalDateTime); canConvert {
			return NewInputLocalDateTime(localDateTimeVal), true
		}

		return nil, false
	}

	// The hook that encodes neo4j durations
	defaultHookDuration EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if duration, canConvert := i.(neo4j.Duration); canConvert {
			return NewInputDuration(&duration), true
		} else if duration, canConvert := i.(*neo4j.Duration); canConvert {
			return NewInputDuration(duration), true
		}

		return nil, false
	}

	// The hook that encodes golang duration values as neo4j durations
	defaultHookGoDuration EncodeHookFunc = func(v reflect.Value, i interface{}) (InputStruct, bool) {
		if duration, canConvert := i.(time.Duration); canConvert {
			return NewInputGoDuration(&duration), true
		} else if duration, canConvert := i.(*time.Duration); canConvert {
			return NewInputGoDuration(duration), true
		}

		return nil, false
	}

	// The hook that encodes neo4j point values
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

	// The hook that encodes neo4j node values
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

	// The hook that encodes neo4j relationship values
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

	// The hook that encodes neo4j path values
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

	// The hook that encodes structs and their exported/tagged fields
	defaultHookStruct = func(tagName string, encoder Encoder) EncodeHookFunc {
		return func(v reflect.Value, i interface{}) (InputStruct, bool) {
			// Verify that the object is a struct
			usedVal := GetValueElem(v)
			if !usedVal.IsValid() || usedVal.Kind() != reflect.Struct {
				return nil, false
			}

			usedType := usedVal.Type()
			resultMap := make(map[string]InputStruct)

			// Run through all of its fields
			for i := 0; i < usedVal.NumField(); i++ {
				field := usedType.Field(i)

				// Get the field tag by the name given in encoder options
				fieldVal := usedVal.FieldByName(field.Name)
				fieldTag := field.Tag.Get(tagName)

				// If there is no tag on the field, skip it
				if fieldTag == "" {
					continue
				}

				// Separate the values of the tag and find its name and omitempty
				allTagValues := strings.Split(fieldTag, ",")
				nameInTag := strings.TrimSpace(allTagValues[0])
				hasOmitEmpty := false
				for _, tagValue := range allTagValues {
					if tagValue == "omitempty" {
						hasOmitEmpty = true
					}
				}

				var key string
				// If the tag name exists and its value exists, then set the resulting map kay as this name. Else, skip this field
				if nameInTag != "" && nameInTag != "-" {
					key = nameInTag
				} else {
					continue
				}

				// If the field is valid and exported, then add it to the resulting map
				// If the field is the zero value of its type, and omitempty was set for it, then skip it
				// If there is a problem, set the mapped value as nil
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
				} else {
					resultMap[key] = nil
				}
			}

			// Every struct will be encoded as maps
			return NewInputMap(resultMap), true
		}
	}

	// The hook that encodes arrays of any type
	defaultHookArray = func(encoder Encoder) EncodeHookFunc {
		return func(v reflect.Value, i interface{}) (InputStruct, bool) {
			usedVal := GetValueElem(v)

			// Only use this hook on array-like object
			if usedVal.Kind() != reflect.Array && usedVal.Kind() != reflect.Slice {
				return nil, false
			}

			// Iterate on every item of the array and try to encode it on its own. Put a nil value if it cannot be encoded
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

	// The hook that encodes arrays of any type
	defaultHookMap = func(encoder Encoder) EncodeHookFunc {
		return func(v reflect.Value, i interface{}) (InputStruct, bool) {
			usedVal := GetValueElem(v)

			if usedVal.Kind() != reflect.Map {
				return nil, false
			}

			encodedMap := make(map[string]InputStruct)

			// Iterate on every key/value of the map
			vMapIter := usedVal.MapRange()
			for vMapIter.Next() {
				key := vMapIter.Key()
				val := vMapIter.Value()

				// Convert the key to string as precisely as we can
				var keyStr string
				switch key.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					keyStr = strconv.FormatInt(key.Int(), 10)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					keyStr = strconv.FormatUint(key.Uint(), 10)
				case reflect.Float32, reflect.Float64:
					keyStr = strconv.FormatFloat(key.Float(), 'f', -1, 64)
				case reflect.Bool:
					keyStr = strconv.FormatBool(key.Bool())
				case reflect.Ptr:
					keyStr = fmt.Sprintf("%v", key.Pointer())
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
