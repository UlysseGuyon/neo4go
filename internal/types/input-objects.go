package types

type InputStruct interface {
	ConvertToMap() map[string]InputStruct
}

type PrimitiveInputObject interface {
	PrimitiveConvert() interface{}
}
