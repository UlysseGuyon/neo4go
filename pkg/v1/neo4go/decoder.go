package neo4go

import (
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

func NewNeo4GoDecoder(options mapstructure.DecoderConfig) (Neo4GoDecoder, internalErr.Neo4GoError) {
	if options.TagName == "" {
		options.TagName = internalMain.DefaultDecodingTagName
	}

	newNeo4GoDecoder := neo4goDecoder{
		DecoderOptions: options,
	}

	return &newNeo4GoDecoder, nil
}

func (decoder *neo4goDecoder) decodeSingleValue(mapInput map[string]interface{}, output interface{}) internalErr.Neo4GoError {
	outputKind := reflect.ValueOf(output).Kind()
	if outputKind != reflect.Ptr && outputKind != reflect.Interface {
		return &internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: "output must ba a pointer",
		}
	}

	decoder.DecoderOptions.Result = output

	mapDecoder, err := mapstructure.NewDecoder(&decoder.DecoderOptions)
	if err != nil {
		return &internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: err.Error(),
		}
	}

	err = mapDecoder.Decode(mapInput)
	if err != nil {
		return &internalErr.Neo4GoUnknownError{}
	}

	return nil
}

func (decoder *neo4goDecoder) DecodeNode(node interface{}, output interface{}) internalErr.Neo4GoError {
	if node == nil {
		return &internalErr.Neo4GoUnknownError{}
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
		return &internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: "Input is not a node or node array",
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
			return &internalErr.Neo4GoUnknownError{}
		}

		err := decoder.decodeSingleValue(usedNode.Props(), output)
		if err != nil {
			return err
		}
	} else {
		for i := 0; i < outputReflect.Len(); i++ {
			if i > len(resultArray) {
				return &internalErr.Neo4GoUnknownError{}
			}
			usedNode := resultArray[i]

			outputReflectItem := outputReflect.Index(i)
			for outputReflectItem.Kind() == reflect.Ptr || outputReflectItem.Kind() == reflect.Interface {
				outputReflectItem = outputReflectItem.Elem()
			}
			if !outputReflectItem.CanInterface() {
				return &internalErr.Neo4GoUnknownError{}
			} else if !outputReflectItem.CanSet() {
				return &internalErr.Neo4GoUnknownError{}
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
	if relationship == nil {
		return &internalErr.Neo4GoUnknownError{}
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
		return &internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: "Input is not a relationship or relationship array",
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
			return &internalErr.Neo4GoUnknownError{}
		}

		err := decoder.decodeSingleValue(usedRelationship.Props(), output)
		if err != nil {
			return err
		}
	} else {
		for i := 0; i < outputReflect.Len(); i++ {
			if i > len(resultArray) {
				return &internalErr.Neo4GoUnknownError{}
			}
			usedRelationship := resultArray[i]

			outputReflectItem := outputReflect.Index(i)
			for outputReflectItem.Kind() == reflect.Ptr || outputReflectItem.Kind() == reflect.Interface {
				outputReflectItem = outputReflectItem.Elem()
			}
			if !outputReflectItem.CanInterface() {
				return &internalErr.Neo4GoUnknownError{}
			} else if !outputReflectItem.CanSet() {
				return &internalErr.Neo4GoUnknownError{}
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
	if path == nil {
		return &internalErr.Neo4GoUnknownError{}
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
		return &internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: "Input is not a path",
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
