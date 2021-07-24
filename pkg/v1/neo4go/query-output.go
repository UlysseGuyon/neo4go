package neo4go

import (
	"fmt"
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
	Nodes     map[string]neo4j.Node
	Relations map[string]neo4j.Relationship
	Paths     map[string]neo4j.Path
	Others    map[string]interface{}
}

// newEmptyRecordMap returns a new RecordMap with each field initialized as an empty map
func newEmptyRecordMap() RecordMap {
	return RecordMap{
		Arrays:    make(map[string]RecordArray),
		Maps:      make(map[string]RecordMap),
		Strings:   make(map[string]string),
		Ints:      make(map[string]int64),
		Floats:    make(map[string]float64),
		Bools:     make(map[string]bool),
		Times:     make(map[string]time.Time),
		Nodes:     make(map[string]neo4j.Node),
		Relations: make(map[string]neo4j.Relationship),
		Paths:     make(map[string]neo4j.Path),
		Others:    make(map[string]interface{}),
	}
}

// RecordMap contains all the typed objects retrieved from a neo4j array in a result.
// The items in the array can be iterated through via the Next function and retrieved as typed objects via the good function
type RecordArray interface {
	// Next iterates to the next item of the array and return false if there are no more items
	Next() bool

	// CurrentAsArray returns the current item of the iteration typed as an Array.
	// The second result is false if the current item cannot be converted as an Array.
	CurrentAsArray() (RecordArray, Neo4GoError)

	// CurrentAsMap returns the current item of the iteration typed as a Map.
	// The second result is false if the current item cannot be converted as a Map.
	CurrentAsMap() (*RecordMap, Neo4GoError)

	// CurrentAsString returns the current item of the iteration typed as a String.
	// The second result is false if the current item cannot be converted as a String.
	CurrentAsString() (*string, Neo4GoError)

	// CurrentAsInt returns the current item of the iteration typed as an Int.
	// The second result is false if the current item cannot be converted as an Int.
	CurrentAsInt() (*int64, Neo4GoError)

	// CurrentAsFloat returns the current item of the iteration typed as a Float.
	// The second result is false if the current item cannot be converted as a Float.
	CurrentAsFloat() (*float64, Neo4GoError)

	// CurrentAsBool returns the current item of the iteration typed as a Bool.
	// The second result is false if the current item cannot be converted as a Bool.
	CurrentAsBool() (*bool, Neo4GoError)

	// CurrentAsTime returns the current item of the iteration typed as a Time.
	// The second result is false if the current item cannot be converted as a Time.
	CurrentAsTime() (*time.Time, Neo4GoError)

	// CurrentAsNode returns the current item of the iteration typed as a Node.
	// The second result is false if the current item cannot be converted as a Node.
	CurrentAsNode() (neo4j.Node, Neo4GoError)

	// CurrentAsRelation returns the current item of the iteration typed as a Relation.
	// The second result is false if the current item cannot be converted as a Relation.
	CurrentAsRelation() (neo4j.Relationship, Neo4GoError)

	// CurrentAsPath returns the current item of the iteration typed as a Path.
	// The second result is false if the current item cannot be converted as a Path.
	CurrentAsPath() (neo4j.Path, Neo4GoError)

	// CurrentAsInterface returns the current item of the iteration typed as an untyped Interface.
	CurrentAsInterface() interface{}

	// TODO comment
	CollectAsArrays() ([]RecordArray, Neo4GoError)
	CollectAsMaps() ([]RecordMap, Neo4GoError)
	CollectAsStrings() ([]string, Neo4GoError)
	CollectAsInts() ([]int64, Neo4GoError)
	CollectAsFloats() ([]float64, Neo4GoError)
	CollectAsBools() ([]bool, Neo4GoError)
	CollectAsTimes() ([]time.Time, Neo4GoError)
	CollectAsNodes() ([]neo4j.Node, Neo4GoError)
	CollectAsRelations() ([]neo4j.Relationship, Neo4GoError)
	CollectAsPaths() ([]neo4j.Path, Neo4GoError)
	CollectAsInterfaces() []interface{}
}

// recordArray is the default implementation of RecordArray
type recordArray struct {
	// The array of untyped objects
	rawArray []interface{}

	// The current index of the iteration in the rawArray
	currentIndex int

	// Tells if the Next function was called at least once on this array
	firstNext bool
}

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

func (rec *recordArray) CollectAsMaps() ([]RecordMap, Neo4GoError) {
	resultArray := make([]RecordMap, 0, len(rec.rawArray))

	for rec.Next() {
		convertedItem, err := rec.CurrentAsMap()
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, *convertedItem)
	}

	return resultArray, nil
}

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

func (rec *recordArray) CollectAsInterfaces() []interface{} {
	return rec.rawArray
}

// NewRecordArray creates a new instance of RecordArray, with a given interface array.
func NewRecordArray(rawArray []interface{}) RecordArray {
	return &recordArray{rawArray: rawArray, currentIndex: 0, firstNext: true}
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
	if rec.currentIndex+1 >= len(rec.rawArray) {
		return false
	}

	if rec.firstNext {
		rec.firstNext = false
	} else {
		rec.currentIndex++
	}

	return true
}

// CurrentAsArray returns the current item of the iteration typed as an Array.
// The second result is false if the current item cannot be converted as an Array.
func (rec *recordArray) CurrentAsArray() (RecordArray, Neo4GoError) {
	if arrayInterface, canConvert := rec.getCurrent().([]interface{}); canConvert {
		return NewRecordArray(arrayInterface), nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into Array",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"[]interface{}"},
	}
}

// CurrentAsMap returns the current item of the iteration typed as a Map.
// The second result is false if the current item cannot be converted as a Map.
func (rec *recordArray) CurrentAsMap() (*RecordMap, Neo4GoError) {
	if mapInterface, canConvert := rec.getCurrent().(map[string]interface{}); canConvert {
		recordMap := decodeMap(mapInterface)
		return &recordMap, nil
	}

	return nil, &internalErr.TypeError{
		Err:           "Could not convert current item of RecordMap into map",
		GotType:       fmt.Sprintf("%T", rec.getCurrent()),
		ExpectedTypes: []string{"map[string]interface{}"},
	}
}

// CurrentAsString returns the current item of the iteration typed as a String.
// The second result is false if the current item cannot be converted as a String.
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
// The second result is false if the current item cannot be converted as an Int.
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
// The second result is false if the current item cannot be converted as a Float.
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
// The second result is false if the current item cannot be converted as a Bool.
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
// The second result is false if the current item cannot be converted as a Time.
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
// The second result is false if the current item cannot be converted as a Node.
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
// The second result is false if the current item cannot be converted as a Relation.
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
// The second result is false if the current item cannot be converted as a Path.
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

// QueryResult is the equivalent of neo4j.Result for neo4go manager queries
type QueryResult interface {
	// Keys returns the keys available on the result set.
	Keys() ([]string, error)

	// Next returns true only if there is a record to be processed.
	Next() bool

	// Err returns the latest error that caused this Next to return false.
	Err() error

	// Record returns the current typed record.
	Record() (*RecordMap, Neo4GoError)

	// Summary returns the summary information about the statement execution.
	Summary() (neo4j.ResultSummary, error)

	// Consume consumes the entire result and returns the summary information
	// about the statement execution.
	Consume() (neo4j.ResultSummary, error)

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

		return nil, ToDriverError(err)
	}

	if from.Next() {
		record, err = from.Record()
		if err != nil {
			return nil, ToDriverError(err)
		}
	}

	if err := from.Err(); err != nil {
		return nil, ToDriverError(err)
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
		return nil, ToDriverError(err)
	}

	for from.Next() {
		record, err := from.Record()
		if err != nil {
			return nil, ToDriverError(err)
		}
		if record == nil {
			return nil, &internalErr.QueryError{
				Err: "Result contains a null record",
			}
		}
		list = append(list, *record)
	}

	if err := from.Err(); err != nil {
		return nil, ToDriverError(err)
	}

	return list, nil
}

// decodeItemInRecordMap takes a string key and an interface value and put it inside the right map field of the result
func decodeItemInRecordMap(key string, value interface{}, resultRecord *RecordMap) {
	switch typedVal := value.(type) {
	case []interface{}:
		resultRecord.Arrays[key] = NewRecordArray(typedVal)
	case map[string]interface{}:
		innerRecordMap := decodeMap(typedVal)
		resultRecord.Maps[key] = innerRecordMap
	case string:
		resultRecord.Strings[key] = typedVal
	case int64:
		resultRecord.Ints[key] = typedVal
	case float64:
		resultRecord.Floats[key] = typedVal
	case bool:
		resultRecord.Bools[key] = typedVal
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
func decodeMap(mapInterface map[string]interface{}) RecordMap {
	newRecordMap := newEmptyRecordMap()

	for key, val := range mapInterface {
		if val == nil {
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
}

// newQueryResult creates a new instance of QueryResult, with a given raw neo4j query result.
func newQueryResult(result neo4j.Result) QueryResult {
	return &queryResult{result: result}
}

// Keys returns the keys available on the result set.
func (res *queryResult) Keys() ([]string, error) {
	return res.result.Keys()
}

// Next returns true only if there is a record to be processed.
func (res *queryResult) Next() bool {
	return res.result.Next()
}

// Err returns the latest error that caused this Next to return false.
func (res *queryResult) Err() error {
	return res.result.Err()
}

// Record returns the current typed record.
func (res *queryResult) Record() (*RecordMap, Neo4GoError) {
	record := res.result.Record()

	newRecordMap := newEmptyRecordMap()
	for _, recordKey := range record.Keys() {
		recordValue, exists := record.Get(recordKey)
		if !exists || recordValue == nil {
			continue
		}

		decodeItemInRecordMap(recordKey, recordValue, &newRecordMap)
	}

	return &newRecordMap, nil
}

// Summary returns the summary information about the statement execution.
func (res *queryResult) Summary() (neo4j.ResultSummary, error) {
	return res.result.Summary()
}

// Consume consumes the entire result and returns the summary information
// about the statement execution.
func (res *queryResult) Consume() (neo4j.ResultSummary, error) {
	return res.result.Consume()
}

// RawResult allow to retrieve the non-typed neo4j.Result
func (res *queryResult) RawResult() neo4j.Result {
	return res.result
}
