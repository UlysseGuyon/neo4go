package neo4go

import (
	"reflect"
	"testing"
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func Test_convertInputObject(t *testing.T) {
	type args struct {
		obj InputStruct
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertInputObject(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputArray(t *testing.T) {
	type args struct {
		value []InputStruct
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputArray(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputArray_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputArray
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputArray.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputArray_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputArray
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputArray.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputArray_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputArray
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputArray.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputMap(t *testing.T) {
	type args struct {
		value map[string]InputStruct
	}
	tests := []struct {
		name string
		args args
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputMap(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputMap_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputMap
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputMap.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputInteger(t *testing.T) {
	type args struct {
		value *int64
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputInteger(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputInteger_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputInteger
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputInteger.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputInteger_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputInteger
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputInteger.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputInteger_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputInteger
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputInteger.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputFloat(t *testing.T) {
	type args struct {
		value *float64
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputFloat(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputFloat_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputFloat
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputFloat.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputFloat_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputFloat
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputFloat.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputFloat_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputFloat
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputFloat.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputBool(t *testing.T) {
	type args struct {
		value *bool
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputBool(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputBool_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputBool
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputBool.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputBool_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputBool
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputBool.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputBool_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputBool
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputBool.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputString(t *testing.T) {
	type args struct {
		value *string
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputString_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputString
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputString.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputString_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputString
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputString.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputString_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputString
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputString.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputByteArray(t *testing.T) {
	type args struct {
		value []byte
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputByteArray(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputByteArray_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputByteArray
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputByteArray.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputByteArray_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputByteArray
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputByteArray.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputByteArray_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputByteArray
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputByteArray.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputDate(t *testing.T) {
	type args struct {
		value *neo4j.Date
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputDate(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDate_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDate
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDate.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDate_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDate
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDate.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDate_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDate
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDate.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputTime(t *testing.T) {
	type args struct {
		value *neo4j.OffsetTime
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputTime(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputTime_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputTime
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputTime.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputTime_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputTime
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputTime.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputTime_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputTime
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputTime.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputLocalTime(t *testing.T) {
	type args struct {
		value *neo4j.LocalTime
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputLocalTime(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputLocalTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputLocalTime_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputLocalTime
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputLocalTime.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputLocalTime_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputLocalTime
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputLocalTime.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputLocalTime_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputLocalTime
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputLocalTime.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputDateTime(t *testing.T) {
	type args struct {
		value *time.Time
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputDateTime(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDateTime_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDateTime
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDateTime.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDateTime_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDateTime
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDateTime.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDateTime_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDateTime
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDateTime.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputLocalDateTime_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputLocalDateTime
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputLocalDateTime.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputLocalDateTime_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputLocalDateTime
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputLocalDateTime.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputLocalDateTime(t *testing.T) {
	type args struct {
		value *neo4j.LocalDateTime
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputLocalDateTime(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputLocalDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputLocalDateTime_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputLocalDateTime
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputLocalDateTime.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDuration_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDuration
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDuration.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDuration_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDuration
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDuration.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputDuration(t *testing.T) {
	type args struct {
		value *neo4j.Duration
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputDuration(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputDuration_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputDuration
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputDuration.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputPoint(t *testing.T) {
	type args struct {
		value *neo4j.Point
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputPoint(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputPoint_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputPoint
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputPoint.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputPoint_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputPoint
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputPoint.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputPoint_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputPoint
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputPoint.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputNode(t *testing.T) {
	type args struct {
		value neo4j.Node
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputNode(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputNode_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputNode
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputNode.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputNode_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputNode
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputNode.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputNode_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputNode
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputNode.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputRelationship(t *testing.T) {
	type args struct {
		value neo4j.Relationship
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputRelationship(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputRelationship() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputRelationship_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputRelationship
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputRelationship.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputRelationship_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputRelationship
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputRelationship.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputRelationship_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputRelationship
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputRelationship.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputPath_ConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		val  *inputPath
		want map[string]InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputPath.ConvertToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputPath_ConvertToInputObject(t *testing.T) {
	tests := []struct {
		name string
		val  *inputPath
		want InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.ConvertToInputObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputPath.ConvertToInputObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInputPath(t *testing.T) {
	type args struct {
		value neo4j.Path
	}
	tests := []struct {
		name string
		args args
		want primitiveInputObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInputPath(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInputPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inputPath_PrimitiveConvert(t *testing.T) {
	tests := []struct {
		name string
		val  *inputPath
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.val.PrimitiveConvert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputPath.PrimitiveConvert() = %v, want %v", got, tt.want)
			}
		})
	}
}
