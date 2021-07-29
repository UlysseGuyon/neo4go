package neo4go

import (
	"fmt"
	"reflect"
	"time"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// RecordMap contains all the typed objects retrieved from a neo4j query result.
// Each object is in its typed map under the key attributed to it inside the cypher query
type RecordMap struct {
	Arrays    map[string]RecordArray
	Maps      map[string]RecordMap
	Strings   map[string]string
	Ints      map[string]int64
	Floats    map[string]float64
	Bools     map[string]bool
	Times     map[string]time.Time
	Durations map[string]time.Duration
	Nodes     map[string]neo4j.Node
	Relations map[string]neo4j.Relationship
	Paths     map[string]neo4j.Path
	Others    map[string]interface{}

	flags queryOutputFlag
}

// newEmptyRecordMap returns a new RecordMap with each field initialized as an empty map
func newEmptyRecordMap(flags queryOutputFlag) RecordMap {
	return RecordMap{
		Arrays:    make(map[string]RecordArray),
		Maps:      make(map[string]RecordMap),
		Strings:   make(map[string]string),
		Ints:      make(map[string]int64),
		Floats:    make(map[string]float64),
		Bools:     make(map[string]bool),
		Times:     make(map[string]time.Time),
		Durations: make(map[string]time.Duration),
		Nodes:     make(map[string]neo4j.Node),
		Relations: make(map[string]neo4j.Relationship),
		Paths:     make(map[string]neo4j.Path),
		Others:    make(map[string]interface{}),
		flags:     flags,
	}
}

// IsEmpty tells if none of the maps in the record have any value
func (rec *RecordMap) IsEmpty() bool {
	return len(rec.Arrays) == 0 && len(rec.Maps) == 0 && len(rec.Strings) == 0 &&
		len(rec.Ints) == 0 && len(rec.Floats) == 0 && len(rec.Bools) == 0 &&
		len(rec.Times) == 0 && len(rec.Durations) == 0 && len(rec.Nodes) == 0 &&
		len(rec.Relations) == 0 && len(rec.Paths) == 0 && len(rec.Others) == 0
}

// RawMap returns this RecordMap as a plain map[string]interface{}
func (rec *RecordMap) RawMap() map[string]interface{} {
	resMap := make(map[string]interface{})

	for key, val := range rec.Arrays {
		resMap[key] = val.CollectAsInterfaces()
	}
	for key, val := range rec.Maps {
		resMap[key] = val.RawMap()
	}
	for key, val := range rec.Strings {
		resMap[key] = val
	}
	for key, val := range rec.Ints {
		resMap[key] = val
	}
	for key, val := range rec.Floats {
		resMap[key] = val
	}
	for key, val := range rec.Bools {
		resMap[key] = val
	}
	for key, val := range rec.Times {
		resMap[key] = val
	}
	for key, val := range rec.Durations {
		resMap[key] = val
	}
	for key, val := range rec.Nodes {
		resMap[key] = val.Props()
	}
	for key, val := range rec.Relations {
		resMap[key] = val.Props()
	}
	// NOTE we currently do not include any path in the map
	for key, val := range rec.Others {
		resMap[key] = val
	}

	return resMap
}

// DecodeNode is an utilitary function that automatically decodes a node from the record object
func (rec *RecordMap) DecodeNode(decoder Decoder, nodeName string, outpout interface{}) Neo4GoError {
	node, exists := rec.Nodes[nodeName]
	if !exists {
		return &internalErr.QueryError{Err: fmt.Sprintf("Node '%s' was not found in record", nodeName)}
	}

	if decoder == nil {
		decoder = NewDecoder(nil)
	}

	return decoder.DecodeNode(&node, outpout)
}

// DecodeRelation is an utilitary function that automatically decodes a relationship from the record object
func (rec *RecordMap) DecodeRelation(decoder Decoder, relationName string, outpout interface{}) Neo4GoError {
	relation, exists := rec.Relations[relationName]
	if !exists {
		return &internalErr.QueryError{Err: fmt.Sprintf("Relationship '%s' was not found in record", relationName)}
	}

	if decoder == nil {
		decoder = NewDecoder(nil)
	}

	return decoder.DecodeRelationship(&relation, outpout)
}

// DecodePath is an utilitary function that automatically decodes a path from the record object
func (rec *RecordMap) DecodePath(decoder Decoder, pathName string, outputNode interface{}, outputRelation interface{}) Neo4GoError {
	path, exists := rec.Paths[pathName]
	if !exists {
		return &internalErr.QueryError{Err: fmt.Sprintf("Path '%s' was not found in record", pathName)}
	}

	if decoder == nil {
		decoder = NewDecoder(nil)
	}

	return decoder.DecodePath(&path, outputNode, outputRelation)
}

// RecordMap contains all the typed objects retrieved from a neo4j array in a result.
// The items in the array can be iterated through via the Next function and retrieved as typed objects via the good function
type RecordArray interface {
	// Next iterates to the next item of the array and return false if there are no more items
	Next() bool

	// Len returns the length of this array
	Len() int

	// CurrentAsArray returns the current item of the iteration typed as an Array.
	// The second result is a non-nil error if the current item cannot be converted as an Array.
	CurrentAsArray() (RecordArray, Neo4GoError)

	// CurrentAsMap returns the current item of the iteration typed as a Map.
	// The second result is a non-nil error if the current item cannot be converted as a Map.
	CurrentAsMap() (*RecordMap, Neo4GoError)

	// CurrentAsString returns the current item of the iteration typed as a String.
	// The second result is a non-nil error if the current item cannot be converted as a String.
	CurrentAsString() (*string, Neo4GoError)

	// CurrentAsInt returns the current item of the iteration typed as an Int.
	// The second result is a non-nil error if the current item cannot be converted as an Int.
	CurrentAsInt() (*int64, Neo4GoError)

	// CurrentAsFloat returns the current item of the iteration typed as a Float.
	// The second result is a non-nil error if the current item cannot be converted as a Float.
	CurrentAsFloat() (*float64, Neo4GoError)

	// CurrentAsBool returns the current item of the iteration typed as a Bool.
	// The second result is a non-nil error if the current item cannot be converted as a Bool.
	CurrentAsBool() (*bool, Neo4GoError)

	// CurrentAsTime returns the current item of the iteration typed as a Time.
	// The second result is a non-nil error if the current item cannot be converted as a Time.
	CurrentAsTime() (*time.Time, Neo4GoError)

	// CurrentAsNode returns the current item of the iteration typed as a Node.
	// The second result is a non-nil error if the current item cannot be converted as a Node.
	CurrentAsNode() (neo4j.Node, Neo4GoError)

	// CurrentAsRelation returns the current item of the iteration typed as a Relation.
	// The second result is a non-nil error if the current item cannot be converted as a Relation.
	CurrentAsRelation() (neo4j.Relationship, Neo4GoError)

	// CurrentAsPath returns the current item of the iteration typed as a Path.
	// The second result is a non-nil error if the current item cannot be converted as a Path.
	CurrentAsPath() (neo4j.Path, Neo4GoError)

	// CurrentAsInterface returns the current item of the iteration typed as an untyped Interface.
	CurrentAsInterface() interface{}

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of RecordArray.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as an Array.
	CollectAsArrays() ([]RecordArray, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of RecordMap.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a Map.
	CollectAsMaps() ([]RecordMap, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of string.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a string.
	CollectAsStrings() ([]string, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of int64.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as an int64.
	CollectAsInts() ([]int64, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of float.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a float.
	CollectAsFloats() ([]float64, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of bool.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a bool.
	CollectAsBools() ([]bool, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of time object.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a time.Time.
	CollectAsTimes() ([]time.Time, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of nodes.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a neo4j.Node.
	CollectAsNodes() ([]neo4j.Node, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of relationships.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a neo4j.Relationship.
	CollectAsRelations() ([]neo4j.Relationship, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of paths.
	// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a neo4j.Path.
	CollectAsPaths() ([]neo4j.Path, Neo4GoError)

	// CollectAsArrays returns the whole array of this RecordArray typed as an Array of interfaces.
	CollectAsInterfaces() []interface{}

	// CollectAndDecodeAsNodes collects all items of this array and converts them as nodes, then decodes them into the output interface
	CollectAndDecodeAsNodes(Decoder, interface{}) Neo4GoError

	// CollectAndDecodeAsRelations collects all items of this array and converts them as relationships, then decodes them into the output interface
	CollectAndDecodeAsRelations(Decoder, interface{}) Neo4GoError
}

// recordArray is the default implementation of RecordArray
type recordArray struct {
	// The array of untyped objects
	rawArray []interface{}

	// The current index of the iteration in the rawArray
	currentIndex int

	// Tells if the Next function was called at least once on this array
	firstNext bool

	flags queryOutputFlag
}

// NewRecordArray creates a new instance of RecordArray, with a given interface array.
func NewRecordArray(rawArray []interface{}, flags queryOutputFlag) RecordArray {
	return &recordArray{rawArray: rawArray, currentIndex: 0, firstNext: true, flags: flags}
}

// getCurrent returns the item currently pointed to in the array, or nil if the index is out of bounds
func (rec *recordArray) getCurrent() interface{} {
	if rec.currentIndex < 0 || rec.currentIndex >= len(rec.rawArray) {
		return nil
	} else {
		return rec.rawArray[rec.currentIndex]
	}
}

// Next iterates to the next item of the array and return false if there are no more items
func (rec *recordArray) Next() bool {
	if rec.firstNext {
		rec.firstNext = false
	} else if rec.currentIndex < len(rec.rawArray) {
		rec.currentIndex++
	}

	if rec.currentIndex >= len(rec.rawArray) {
		return false
	}

	return true
}

// Len returns the length of this array
func (rec *recordArray) Len() int {
	return len(rec.rawArray)
}

// CurrentAsArray returns the current item of the iteration typed as an Array.
// The second result is a non-nil error if the current item cannot be converted as an Array.
func (rec *recordArray) CurrentAsArray() (RecordArray, Neo4GoError) {
	if arrayInterface, canConvert := rec.getCurrent().([]interface{}); canConvert {
		return NewRecordArray(arrayInterface, rec.flags), nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into Array",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"[]interface{}"},
	}
}

// CurrentAsMap returns the current item of the iteration typed as a Map.
// The second result is a non-nil error if the current item cannot be converted as a Map.
func (rec *recordArray) CurrentAsMap() (*RecordMap, Neo4GoError) {
	if mapInterface, canConvert := rec.getCurrent().(map[string]interface{}); canConvert {
		recordMap := decodeMap(mapInterface, rec.flags)
		return &recordMap, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into map",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"map[string]interface{}"},
	}
}

// CurrentAsString returns the current item of the iteration typed as a String.
// The second result is a non-nil error if the current item cannot be converted as a String.
func (rec *recordArray) CurrentAsString() (*string, Neo4GoError) {
	if converted, canConvert := rec.getCurrent().(string); canConvert {
		return &converted, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into string",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"string"},
	}
}

// CurrentAsInt returns the current item of the iteration typed as an Int.
// The second result is a non-nil error if the current item cannot be converted as an Int.
func (rec *recordArray) CurrentAsInt() (*int64, Neo4GoError) {
	if converted, canConvert := rec.getCurrent().(int64); canConvert {
		return &converted, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into int",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"int64"},
	}
}

// CurrentAsFloat returns the current item of the iteration typed as a Float.
// The second result is a non-nil error if the current item cannot be converted as a Float.
func (rec *recordArray) CurrentAsFloat() (*float64, Neo4GoError) {
	if converted, canConvert := rec.getCurrent().(float64); canConvert {
		return &converted, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into float",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"float64"},
	}
}

// CurrentAsBool returns the current item of the iteration typed as a Bool.
// The second result is a non-nil error if the current item cannot be converted as a Bool.
func (rec *recordArray) CurrentAsBool() (*bool, Neo4GoError) {
	if converted, canConvert := rec.getCurrent().(bool); canConvert {
		return &converted, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into bool",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"bool"},
	}
}

// CurrentAsTime returns the current item of the iteration typed as a Time.
// The second result is a non-nil error if the current item cannot be converted as a Time.
func (rec *recordArray) CurrentAsTime() (*time.Time, Neo4GoError) {
	if converted, canConvert := rec.getCurrent().(time.Time); canConvert {
		return &converted, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into time object",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"time.Time"},
	}
}

// CurrentAsNode returns the current item of the iteration typed as a Node.
// The second result is a non-nil error if the current item cannot be converted as a Node.
func (rec *recordArray) CurrentAsNode() (neo4j.Node, Neo4GoError) {
	if converted, canConvert := rec.getCurrent().(neo4j.Node); canConvert {
		return converted, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into node",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"neo4j.Node"},
	}
}

// CurrentAsRelation returns the current item of the iteration typed as a Relation.
// The second result is a non-nil error if the current item cannot be converted as a Relation.
func (rec *recordArray) CurrentAsRelation() (neo4j.Relationship, Neo4GoError) {
	if converted, canConvert := rec.getCurrent().(neo4j.Relationship); canConvert {
		return converted, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into relationship",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"neo4j.Relationship"},
	}
}

// CurrentAsPath returns the current item of the iteration typed as a Path.
// The second result is a non-nil error if the current item cannot be converted as a Path.
func (rec *recordArray) CurrentAsPath() (neo4j.Path, Neo4GoError) {
	if converted, canConvert := rec.getCurrent().(neo4j.Path); canConvert {
		return converted, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into path",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"neo4j.Path"},
	}
}

// CurrentAsInterface returns the current item of the iteration typed as an untyped Interface.
func (rec *recordArray) CurrentAsInterface() interface{} {
	return rec.getCurrent()
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of RecordArray.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as an Array.
func (rec *recordArray) CollectAsArrays() ([]RecordArray, Neo4GoError) {
	resultArray := make([]RecordArray, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsArray()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of RecordMap.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a Map.
func (rec *recordArray) CollectAsMaps() ([]RecordMap, Neo4GoError) {
	resultArray := make([]RecordMap, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsMap()
		if err != nil {
			return nil, err
		}
		if !convertedItem.IsEmpty() || rec.flags.HasBaseQueryOuputFlag(KEEP_EMPTY_MAPS) {
			resultArray = append(resultArray, *convertedItem)
		}
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of string.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a string.
func (rec *recordArray) CollectAsStrings() ([]string, Neo4GoError) {
	resultArray := make([]string, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsString()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, *convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of int64.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as an int64.
func (rec *recordArray) CollectAsInts() ([]int64, Neo4GoError) {
	resultArray := make([]int64, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsInt()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, *convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of float.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a float.
func (rec *recordArray) CollectAsFloats() ([]float64, Neo4GoError) {
	resultArray := make([]float64, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsFloat()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, *convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of bool.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a bool.
func (rec *recordArray) CollectAsBools() ([]bool, Neo4GoError) {
	resultArray := make([]bool, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsBool()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, *convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of time object.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a time.Time.
func (rec *recordArray) CollectAsTimes() ([]time.Time, Neo4GoError) {
	resultArray := make([]time.Time, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsTime()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, *convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of nodes.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a neo4j.Node.
func (rec *recordArray) CollectAsNodes() ([]neo4j.Node, Neo4GoError) {
	resultArray := make([]neo4j.Node, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsNode()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of relationships.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a neo4j.Relationship.
func (rec *recordArray) CollectAsRelations() ([]neo4j.Relationship, Neo4GoError) {
	resultArray := make([]neo4j.Relationship, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsRelation()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of paths.
// The second result is a non-nil error if at least one item of the RecordArray cannot be converted as a neo4j.Path.
func (rec *recordArray) CollectAsPaths() ([]neo4j.Path, Neo4GoError) {
	resultArray := make([]neo4j.Path, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsPath()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, convertedItem)
	}

	return resultArray, nil
}

// CollectAsArrays returns the whole array of this RecordArray typed as an Array of interfaces.
func (rec *recordArray) CollectAsInterfaces() []interface{} {
	return rec.rawArray
}

// CollectAndDecodeAsNodes collects all items of this array and converts them as nodes, then decodes them into the output interface
func (rec *recordArray) CollectAndDecodeAsNodes(decoder Decoder, output interface{}) Neo4GoError {
	if decoder == nil {
		decoder = NewDecoder(nil)
	}

	nodeList, err := rec.CollectAsNodes()
	if err != nil {
		return err
	}

	return decoder.DecodeNode(&nodeList, output)
}

// CollectAndDecodeAsRelations collects all items of this array and converts them as relationships, then decodes them into the output interface
func (rec *recordArray) CollectAndDecodeAsRelations(decoder Decoder, output interface{}) Neo4GoError {
	if decoder == nil {
		decoder = NewDecoder(nil)
	}

	relationList, err := rec.CollectAsRelations()
	if err != nil {
		return err
	}

	return decoder.DecodeRelationship(&relationList, output)
}

// QueryResult is the equivalent of neo4j.Result for neo4go manager queries
type QueryResult interface {
	// Keys returns the keys available on the result set.
	Keys() ([]string, Neo4GoError)

	// Next returns true only if there is a record to be processed.
	Next() bool

	// Err returns the latest error that caused this Next to return false.
	Err() Neo4GoError

	// Record returns the current typed record.
	Record() (*RecordMap, Neo4GoError)

	// Summary returns the summary information about the statement execution.
	Summary() (neo4j.ResultSummary, Neo4GoError)

	// Consume consumes the entire result and returns the summary information
	// about the statement execution.
	Consume() (neo4j.ResultSummary, Neo4GoError)

	// RawResult allow to retrieve the non-typed neo4j.Result
	RawResult() neo4j.Result
}

// Single returns one and only one record from the result stream. Any error passed in
// or reported while navigating the result stream is returned without any conversion.
// If the result stream contains zero or more than one records error is returned.
// This function is nearly entirely copied from https://github.com/neo4j/neo4j-go-driver/blob/4.3/neo4j/result_helpers.go
func Single(from QueryResult, err error) (*RecordMap, Neo4GoError) {
	var record *RecordMap

	if err != nil {
		if convertedErr, canConvert := err.(Neo4GoError); canConvert {
			return nil, convertedErr
		}

		return nil, toDriverError(err)
	}

	if from.Next() {
		record, err = from.Record()
		if err != nil {
			return nil, toDriverError(err)
		}
	}

	if err := from.Err(); err != nil {
		return nil, toDriverError(err)
	}

	if record == nil {
		return nil, &internalErr.QueryError{
			Err: "Result contains no record",
		}
	}

	if from.Next() {
		return nil, &internalErr.QueryError{
			Err: "Result contains more than one record",
		}
	}

	return record, nil
}

// Collect loops through the result stream, collects records into a slice and returns the
// resulting slice. Any error passed in or reported while navigating the result stream is
// returned without any conversion.
// This function is nearly entirely copied from https://github.com/neo4j/neo4j-go-driver/blob/4.3/neo4j/result_helpers.go
func Collect(from QueryResult, err error) ([]RecordMap, Neo4GoError) {
	var list []RecordMap

	if err != nil {
		return nil, toDriverError(err)
	}

	for from.Next() {
		record, err := from.Record()
		if err != nil {
			return nil, toDriverError(err)
		}
		if record == nil {
			return nil, &internalErr.QueryError{
				Err: "Result contains a null record",
			}
		}
		list = append(list, *record)
	}

	if err := from.Err(); err != nil {
		return nil, toDriverError(err)
	}

	return list, nil
}

// decodeItemInRecordMap takes a string key and an interface value and put it inside the right map field of the result
func decodeItemInRecordMap(key string, value interface{}, resultRecord *RecordMap) {
	if resultRecord == nil || IsNil(reflect.ValueOf(value)) {
		return
	}

	switch typedVal := value.(type) {
	case []interface{}:
		resultRecord.Arrays[key] = NewRecordArray(typedVal, resultRecord.flags)
	case map[string]interface{}:
		innerRecordMap := decodeMap(typedVal, resultRecord.flags)
		if !innerRecordMap.IsEmpty() || resultRecord.flags.HasBaseQueryOuputFlag(KEEP_EMPTY_MAPS) {
			resultRecord.Maps[key] = innerRecordMap
		}
	case string:
		resultRecord.Strings[key] = typedVal
	case int64:
		resultRecord.Ints[key] = typedVal
	case float64:
		resultRecord.Floats[key] = typedVal
	case bool:
		resultRecord.Bools[key] = typedVal
	case time.Time:
		resultRecord.Times[key] = typedVal
	case neo4j.LocalDateTime:
		resultRecord.Times[key] = typedVal.Time()
	case neo4j.Date:
		resultRecord.Times[key] = typedVal.Time()
	case neo4j.OffsetTime:
		resultRecord.Times[key] = typedVal.Time()
	case neo4j.LocalTime:
		resultRecord.Times[key] = typedVal.Time()
	case neo4j.Duration:
		resultRecord.Durations[key] =
			30*24*time.Hour*time.Duration(typedVal.Months()) +
				24*time.Hour*time.Duration(typedVal.Days()) +
				time.Second*time.Duration(typedVal.Seconds()) +
				time.Nanosecond*time.Duration(typedVal.Nanos())
	case neo4j.Node:
		resultRecord.Nodes[key] = typedVal
	case *neo4j.Node:
		resultRecord.Nodes[key] = *typedVal
	case neo4j.Relationship:
		resultRecord.Relations[key] = typedVal
	case *neo4j.Relationship:
		resultRecord.Relations[key] = *typedVal
	case neo4j.Path:
		resultRecord.Paths[key] = typedVal
	case *neo4j.Path:
		resultRecord.Paths[key] = *typedVal
	default:
		resultRecord.Others[key] = typedVal
	}
}

// decodeMap takes a map as input and converts it into a RecordMap for simplicity of use
func decodeMap(mapInterface map[string]interface{}, flags queryOutputFlag) RecordMap {
	newRecordMap := newEmptyRecordMap(flags)

	for key, val := range mapInterface {
		if IsNil(reflect.ValueOf(val)) && !flags.HasBaseQueryOuputFlag(INCLUDE_NIL_IN_RECORDS) {
			continue
		}

		decodeItemInRecordMap(key, val, &newRecordMap)
	}

	return newRecordMap
}

// queryResult is the default implementation of the QueryResult interface
type queryResult struct {
	// The raw neo4j result of the query
	result neo4j.Result

	// The formatting to apply to the records of this result
	flags queryOutputFlag
}

// newQueryResult creates a new instance of QueryResult, with a given raw neo4j query result.
func newQueryResult(result neo4j.Result, flags queryOutputFlag) QueryResult {
	return &queryResult{result: result, flags: flags}
}

// Keys returns the keys available on the result set.
func (res *queryResult) Keys() ([]string, Neo4GoError) {
	keys, err := res.result.Keys()
	if err != nil {
		return nil, toDriverError(err)
	}

	return keys, nil
}

// Next returns true only if there is a record to be processed.
func (res *queryResult) Next() bool {
	return res.result.Next()
}

// Err returns the latest error that caused this Next to return false.
func (res *queryResult) Err() Neo4GoError {
	err := res.result.Err()
	if err != nil {
		return toDriverError(err)
	}

	return nil
}

// Record returns the current typed record.
func (res *queryResult) Record() (*RecordMap, Neo4GoError) {
	record := res.result.Record()

	newMap := make(map[string]interface{})
	for _, recordKey := range record.Keys() {
		recordValue, exists := record.Get(recordKey)
		if !exists {
			continue
		}
		newMap[recordKey] = recordValue
	}

	resRecord := decodeMap(newMap, res.flags)

	return &resRecord, nil
}

// Summary returns the summary information about the statement execution.
func (res *queryResult) Summary() (neo4j.ResultSummary, Neo4GoError) {
	sum, err := res.result.Summary()
	if err != nil {
		return nil, toDriverError(err)
	}

	return sum, nil
}

// Consume consumes the entire result and returns the summary information
// about the statement execution.
func (res *queryResult) Consume() (neo4j.ResultSummary, Neo4GoError) {
	sum, err := res.result.Consume()
	if err != nil {
		return nil, toDriverError(err)
	}

	return sum, nil
}

// RawResult allow to retrieve the non-typed neo4j.Result
func (res *queryResult) RawResult() neo4j.Result {
	return res.result
}
