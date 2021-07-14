package neo4go

import (
	"fmt"
	"reflect"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	internalMain "github.com/UlysseGuyon/neo4go/internal/neo4go"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Neo4GoDecoder interface {
	DecodeNode(interface{}, interface{}) internalErr.Neo4GoError
	DecodeRelationship(interface{}, interface{}) internalErr.Neo4GoError
	DecodePath(interface{}, interface{}, interface{}) internalErr.Neo4GoError
}

type neo4goDecoder struct {
	DecoderOptions mapstructure.DecoderConfig
}

func NewNeo4GoDecoder(options mapstructure.DecoderConfig) Neo4GoDecoder {
	if options.TagName == "" {
		options.TagName = internalMain.DefaultDecodingTagName
	}

	newNeo4GoDecoder := neo4goDecoder{
		DecoderOptions: options,
	}

	return &newNeo4GoDecoder
}

func (decoder *neo4goDecoder) decodeSingleValue(mapInput map[string]interface{}, output interface{}) internalErr.Neo4GoError {
	outputKind := reflect.ValueOf(output).Kind()
	if outputKind != reflect.Ptr && outputKind != reflect.Interface {
		return &internalErr.TypeError{
			Err:           "Output must be a pointer",
			ExpectedTypes: []string{fmt.Sprintf("%T Pointer", output)},
			GotType:       fmt.Sprintf("%T", output),
		}
	}

	decoder.DecoderOptions.Result = output

	mapDecoder, err := mapstructure.NewDecoder(&decoder.DecoderOptions)
	if err != nil {
		return &internalErr.DecodingError{
			Err: err.Error(),
		}
	}

	err = mapDecoder.Decode(mapInput)
	if err != nil {
		return &internalErr.DecodingError{
			Err: err.Error(),
		}
	}

	return nil
}

func (decoder *neo4goDecoder) DecodeNode(node interface{}, output interface{}) internalErr.Neo4GoError {
	expectedTypes := []string{
		"Node",
		"*Node",
		"[]Node",
		"*[]Node",
		"[]*Node",
		"*[]*Node",
	}

	if node == nil {
		return &internalErr.TypeError{
			Err:           "Decoded node cannot be null",
			ExpectedTypes: expectedTypes,
			GotType:       "null",
		}
	}

	resultArray := make([]neo4j.Node, 0)

	switch typedNode := node.(type) {
	case neo4j.Node:
		resultArray = append(resultArray, typedNode)
	case *neo4j.Node:
		if typedNode != nil {
			resultArray = append(resultArray, *typedNode)
		}
	case []neo4j.Node:
		resultArray = append(resultArray, typedNode...)
	case *[]neo4j.Node:
		if typedNode != nil {
			resultArray = append(resultArray, *typedNode...)
		}
	case []*neo4j.Node:
		for _, n := range typedNode {
			if n != nil {
				resultArray = append(resultArray, *n)
			}
		}
	case *[]*neo4j.Node:
		if typedNode != nil {
			for _, n := range *typedNode {
				if n != nil {
					resultArray = append(resultArray, *n)
				}
			}
		}
	default:
		return &internalErr.TypeError{
			Err:           "Input is not a node or node array",
			ExpectedTypes: expectedTypes,
			GotType:       fmt.Sprintf("%T", node),
		}
	}

	outputReflect := reflect.ValueOf(output)
	for outputReflect.Kind() == reflect.Ptr || outputReflect.Kind() == reflect.Interface {
		outputReflect = outputReflect.Elem()
	}

	isSlice := false
	if outputReflect.Kind() == reflect.Array || outputReflect.Kind() == reflect.Slice {
		isSlice = true
	}

	if !isSlice {
		var usedNode neo4j.Node
		if len(resultArray) > 0 {
			usedNode = resultArray[0]
		} else {
			return &internalErr.DecodingError{
				Err: "Could not decode one node to fit in output",
			}
		}

		err := decoder.decodeSingleValue(usedNode.Props(), output)
		if err != nil {
			return err
		}
	} else {
		for i := 0; i < outputReflect.Len(); i++ {
			if i > len(resultArray) {
				return &internalErr.DecodingError{
					Err: "Could not decode enough nodes to fit in output",
				}
			}
			usedNode := resultArray[i]

			outputReflectItem := outputReflect.Index(i)
			for outputReflectItem.Kind() == reflect.Ptr || outputReflectItem.Kind() == reflect.Interface {
				outputReflectItem = outputReflectItem.Elem()
			}
			if !outputReflectItem.CanInterface() {
				return &internalErr.TypeError{
					Err:           "Output list item cannot be converted as interface",
					ExpectedTypes: []string{"interface{}"},
					GotType:       outputReflectItem.Type().Name(),
				}
			} else if !outputReflectItem.CanSet() {
				return &internalErr.TypeError{
					Err:           "Output list item value cannot be written",
					ExpectedTypes: []string{"Any"},
					GotType:       outputReflectItem.Type().Name(),
				}
			}
			outputItemInterface := outputReflectItem.Interface()

			err := decoder.decodeSingleValue(usedNode.Props(), &outputItemInterface)
			if err != nil {
				return err
			}

			// NOTE Convert should never panic but nothing is impossible :)
			converted := reflect.ValueOf(outputItemInterface).Convert(outputReflectItem.Type())

			outputReflectItem.Set(converted)
		}
	}

	return nil
}

func (decoder *neo4goDecoder) DecodeRelationship(relationship interface{}, output interface{}) internalErr.Neo4GoError {
	expectedTypes := []string{
		"Relationship",
		"*Relationship",
		"[]Relationship",
		"*[]Relationship",
		"[]*Relationship",
		"*[]*Relationship",
	}

	if relationship == nil {
		return &internalErr.TypeError{
			Err:           "Decoded relationship cannot be null",
			ExpectedTypes: expectedTypes,
			GotType:       "null",
		}
	}

	resultArray := make([]neo4j.Relationship, 0)

	switch typedRelationship := relationship.(type) {
	case neo4j.Relationship:
		resultArray = append(resultArray, typedRelationship)
	case *neo4j.Relationship:
		if typedRelationship != nil {
			resultArray = append(resultArray, *typedRelationship)
		}
	case []neo4j.Relationship:
		resultArray = append(resultArray, typedRelationship...)
	case *[]neo4j.Relationship:
		if typedRelationship != nil {
			resultArray = append(resultArray, *typedRelationship...)
		}
	case []*neo4j.Relationship:
		for _, n := range typedRelationship {
			if n != nil {
				resultArray = append(resultArray, *n)
			}
		}
	case *[]*neo4j.Relationship:
		if typedRelationship != nil {
			for _, n := range *typedRelationship {
				if n != nil {
					resultArray = append(resultArray, *n)
				}
			}
		}
	default:
		return &internalErr.TypeError{
			Err:           "Input is not a relationship or relationship array",
			ExpectedTypes: expectedTypes,
			GotType:       fmt.Sprintf("%T", relationship),
		}
	}

	outputReflect := reflect.ValueOf(output)
	for outputReflect.Kind() == reflect.Ptr || outputReflect.Kind() == reflect.Interface {
		outputReflect = outputReflect.Elem()
	}

	isSlice := false
	if outputReflect.Kind() == reflect.Array || outputReflect.Kind() == reflect.Slice {
		isSlice = true
	}

	if !isSlice {
		var usedRelationship neo4j.Relationship
		if len(resultArray) > 0 {
			usedRelationship = resultArray[0]
		} else {
			return &internalErr.DecodingError{
				Err: "Could not decode one relationship to fit in output",
			}
		}

		err := decoder.decodeSingleValue(usedRelationship.Props(), output)
		if err != nil {
			return err
		}
	} else {
		for i := 0; i < outputReflect.Len(); i++ {
			if i > len(resultArray) {
				return &internalErr.DecodingError{
					Err: "Could not decode enough relationships to fit in output",
				}
			}
			usedRelationship := resultArray[i]

			outputReflectItem := outputReflect.Index(i)
			for outputReflectItem.Kind() == reflect.Ptr || outputReflectItem.Kind() == reflect.Interface {
				outputReflectItem = outputReflectItem.Elem()
			}
			if !outputReflectItem.CanInterface() {
				return &internalErr.TypeError{
					Err:           "Output list item cannot be converted as interface",
					ExpectedTypes: []string{"interface{}"},
					GotType:       outputReflectItem.Type().Name(),
				}
			} else if !outputReflectItem.CanSet() {
				return &internalErr.TypeError{
					Err:           "Output list item value cannot be written",
					ExpectedTypes: []string{"Any"},
					GotType:       outputReflectItem.Type().Name(),
				}
			}
			outputItemInterface := outputReflectItem.Interface()

			err := decoder.decodeSingleValue(usedRelationship.Props(), &outputItemInterface)
			if err != nil {
				return err
			}

			// NOTE Convert should never panic but nothing is impossible :)
			converted := reflect.ValueOf(outputItemInterface).Convert(outputReflectItem.Type())

			outputReflectItem.Set(converted)
		}
	}

	return nil
}

func (decoder *neo4goDecoder) DecodePath(path interface{}, outputNodes interface{}, outputRelationships interface{}) internalErr.Neo4GoError {
	expectedTypes := []string{
		"Path",
		"*Path",
	}

	if path == nil {
		return &internalErr.TypeError{
			Err:           "Decoded path cannot be null",
			ExpectedTypes: expectedTypes,
			GotType:       "null",
		}
	}

	resultNodeArray := make([]neo4j.Node, 0)
	resultRelationshipArray := make([]neo4j.Relationship, 0)

	switch typedPath := path.(type) {
	case neo4j.Path:
		resultNodeArray = append(resultNodeArray, typedPath.Nodes()...)
		resultRelationshipArray = append(resultRelationshipArray, typedPath.Relationships()...)
	case *neo4j.Path:
		if typedPath != nil {
			resultNodeArray = append(resultNodeArray, (*typedPath).Nodes()...)
			resultRelationshipArray = append(resultRelationshipArray, (*typedPath).Relationships()...)
		}
	default:
		return &internalErr.TypeError{
			Err:           "Input is not a path",
			ExpectedTypes: expectedTypes,
			GotType:       fmt.Sprintf("%T", path),
		}
	}

	err := decoder.DecodeNode(&resultNodeArray, outputNodes)
	if err != nil {
		return err
	}
	err = decoder.DecodeRelationship(&resultRelationshipArray, outputRelationships)
	if err != nil {
		return err
	}

	return nil
}
