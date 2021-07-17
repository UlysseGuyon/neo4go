package neo4go

import (
	"fmt"
	"reflect"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	internalMain "github.com/UlysseGuyon/neo4go/internal/neo4go"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Decoder allows a user to decode Node, Relationship or Path object from the Neo4J driver
// as custom struct fields, provided the right pointer and that the struct has the right tags on its fields
type Decoder interface {
	// DecodeNode takes a node like object (pointers and lists are accepted) and decodes it in the second argument
	DecodeNode(interface{}, interface{}) internalErr.Neo4GoError

	// DecodeNode takes a relationship like object (pointers and lists are accepted) and decodes it in the second argument
	DecodeRelationship(interface{}, interface{}) internalErr.Neo4GoError

	// DecodeNode takes a path like object (pointers are accepted) and decodes its nodes in the second argument and its relationships in the third
	DecodePath(interface{}, interface{}, interface{}) internalErr.Neo4GoError
}

// neo4goDecoder is the default implementation of the Decoder interface
type neo4goDecoder struct {
	// The options used when decoding an object
	options mapstructure.DecoderConfig
}

// NewDecoder creates a new instance of Decoder, with a given config. A nil config will result in the default config beinng applied
func NewDecoder(options *mapstructure.DecoderConfig) Decoder {
	// Use the given config if not nil
	usedOpt := mapstructure.DecoderConfig{}
	if options != nil {
		usedOpt = *options
	}

	// Use the default decoding tag name if none is given
	if usedOpt.TagName == "" {
		usedOpt.TagName = internalMain.DefaultDecodingTagName
	}

	// Instanciate and return the decoder
	newNeo4GoDecoder := neo4goDecoder{
		options: usedOpt,
	}

	return &newNeo4GoDecoder
}

// decodeSingleValue takes a map of values (typically a Node.Props() or Relationship.Props()) and maps it in the
// outputs fields using the mapstructure package
func (decoder *neo4goDecoder) decodeSingleValue(mapInput map[string]interface{}, output interface{}) internalErr.Neo4GoError {
	// Check that the output is of the right type/kind
	outputKind := reflect.ValueOf(output).Kind()
	if outputKind != reflect.Ptr && outputKind != reflect.Interface {
		return &internalErr.TypeError{
			Err:           "Output must be a pointer",
			ExpectedTypes: []string{fmt.Sprintf("%T Pointer", output)},
			GotType:       fmt.Sprintf("%T", output),
		}
	}

	// Set the mapstructure decoder result as the output
	decoder.options.Result = output

	// Apply the mapstructure decoding after creating a new decoder
	mapDecoder, err := mapstructure.NewDecoder(&decoder.options)
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

// DecodeNode takes a node like object (pointers and lists are accepted) and decodes it in the second argument
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

	// Collect the nodes from the first argument in resultArray
	resultArray := make([]neo4j.Node, 0)

	// Allow plain nodes, pointers, arrays or combinations of them
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

	// Get the reflected value of the output
	outputReflect := GetValueElem(reflect.ValueOf(output))

	// Processing will be different depending on if the output is an array or a plain struct
	isSlice := false
	if outputReflect.Kind() == reflect.Array || outputReflect.Kind() == reflect.Slice {
		isSlice = true
	}

	if !isSlice {
		// If the output is a plain object, we can only decode the first node, even if we had more in the input
		// If there are more than 1 node as input, they will NOT be decoded

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
		// If the output is an array, we try to fill it with as much nodes as we can.
		// If the input does not provide enough nodes, an error is thrown.
		// If the input provides more nodes than the output can stock, the extra ones will NOT be decoded

		for i := 0; i < outputReflect.Len(); i++ {
			if i > len(resultArray) {
				return &internalErr.DecodingError{
					Err: "Could not decode enough nodes to fit in output",
				}
			}
			usedNode := resultArray[i]

			// For each item of the list, get its value and apply the node to its fields it can be set
			outputReflectItem := GetValueElem(outputReflect.Index(i))
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

			// After retreiving the nodes fields inside the interface outputItemInterface, set its value to the reflect value
			// NOTE Convert should never panic but nothing is impossible :)
			converted := reflect.ValueOf(outputItemInterface).Convert(outputReflectItem.Type())

			outputReflectItem.Set(converted)
		}
	}

	return nil
}

// DecodeNode takes a relationship like object (pointers and lists are accepted) and decodes it in the second argument
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

	// Collect the relationships from the first argument in resultArray
	resultArray := make([]neo4j.Relationship, 0)

	// Allow plain relationships, pointers, arrays or combinations of them
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

	// Get the reflected value of the output
	outputReflect := GetValueElem(reflect.ValueOf(output))

	// Processing will be different depending on if the output is an array or a plain struct
	isSlice := false
	if outputReflect.Kind() == reflect.Array || outputReflect.Kind() == reflect.Slice {
		isSlice = true
	}

	if !isSlice {
		// If the output is a plain object, we can only decode the first relationship, even if we had more in the input
		// If there are more than 1 relationship as input, they will NOT be decoded

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
		// If the output is an array, we try to fill it with as much relationships as we can.
		// If the input does not provide enough relationships, an error is thrown.
		// If the input provides more relationships than the output can stock, the extra ones will NOT be decoded

		for i := 0; i < outputReflect.Len(); i++ {
			if i > len(resultArray) {
				return &internalErr.DecodingError{
					Err: "Could not decode enough relationships to fit in output",
				}
			}
			usedRelationship := resultArray[i]

			// For each item of the list, get its value and apply the node to its fields it can be set
			outputReflectItem := GetValueElem(outputReflect.Index(i))
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

			// After retreiving the nodes fields inside the interface outputItemInterface, set its value to the reflect value
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

	// Collect the nodes and relationships from the first argument in these arrays
	resultNodeArray := make([]neo4j.Node, 0)
	resultRelationshipArray := make([]neo4j.Relationship, 0)

	// Allow only path and poiter to path
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

	// Use the existing decoding functions to map the collected nodes and relationships in the second and third arguments (the outputs)
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
