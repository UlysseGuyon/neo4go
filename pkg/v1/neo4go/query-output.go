package neo4go

import (
	"time"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// RecordMap contains all the typed objects retrieved from a neo4j query result.
// Each object is in it typed map under the key attributed to it inside the cypher query
type RecordMap struct {
	Arrays             map[string][]interface{}
	MapArrays          map[string][]RecordMap
	NodeArrays         map[string][]neo4j.Node
	RelationshipArrays map[string][]neo4j.Relationship
	PathArrays         map[string][]neo4j.Path
	Maps               map[string]RecordMap
	Strings            map[string]string
	Ints               map[string]int64
	Floats             map[string]float64
	Bools              map[string]bool
	Times              map[string]time.Time
	Nodes              map[string]neo4j.Node
	Relations          map[string]neo4j.Relationship
	Paths              map[string]neo4j.Path
	Others             map[string]interface{}
}

// newEmptyRecordMap returns a new RecordMap with each field initialized as an empty map
func newEmptyRecordMap() RecordMap {
	return RecordMap{
		Arrays:             make(map[string][]interface{}),
		MapArrays:          make(map[string][]RecordMap),
		NodeArrays:         make(map[string][]neo4j.Node),
		RelationshipArrays: make(map[string][]neo4j.Relationship),
		PathArrays:         make(map[string][]neo4j.Path),
		Maps:               make(map[string]RecordMap),
		Strings:            make(map[string]string),
		Ints:               make(map[string]int64),
		Floats:             make(map[string]float64),
		Bools:              make(map[string]bool),
		Times:              make(map[string]time.Time),
		Nodes:              make(map[string]neo4j.Node),
		Relations:          make(map[string]neo4j.Relationship),
		Paths:              make(map[string]neo4j.Path),
		Others:             make(map[string]interface{}),
	}
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
	Record() (*RecordMap, internalErr.Neo4GoError)

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
func Single(from QueryResult, err error) (*RecordMap, internalErr.Neo4GoError) {
	var record *RecordMap

	if err != nil {
		if convertedErr, canConvert := err.(internalErr.Neo4GoError); canConvert {
			return nil, convertedErr
		}

		return nil, internalErr.ToDriverError(err)
	}

	if from.Next() {
		record, err = from.Record()
		if err != nil {
			return nil, internalErr.ToDriverError(err)
		}
	}

	if err := from.Err(); err != nil {
		return nil, internalErr.ToDriverError(err)
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
func Collect(from QueryResult, err error) ([]RecordMap, internalErr.Neo4GoError) {
	var list []RecordMap

	if err != nil {
		return nil, internalErr.ToDriverError(err)
	}

	for from.Next() {
		record, err := from.Record()
		if err != nil {
			return nil, internalErr.ToDriverError(err)
		}
		if record == nil {
			return nil, &internalErr.QueryError{
				Err: "Result contains a null record",
			}
		}
		list = append(list, *record)
	}

	if err := from.Err(); err != nil {
		return nil, internalErr.ToDriverError(err)
	}

	return list, nil
}

// decodeItemInRecordMap takes a string key and an interface value and put it inside the right map field of the result
func decodeItemInRecordMap(key string, value interface{}, resultRecord *RecordMap) internalErr.Neo4GoError {
	switch typedVal := value.(type) {
	case []interface{}:
		// If the value is an array, first check if the items can be typed, else return an interface array

		// Check maps array
		newRecordMapArray := make([]RecordMap, 0)
		for _, item := range typedVal {
			if convertedItem, canConvert := item.(map[string]interface{}); canConvert {
				innerRecordMap, err := decodeMap(convertedItem)
				if err != nil {
					return err
				}
				newRecordMapArray = append(newRecordMapArray, *innerRecordMap)
			}
		}

		// Check nodes array
		newNodeArray := make([]neo4j.Node, 0)
		for _, item := range typedVal {
			if convertedItem, canConvert := item.(neo4j.Node); canConvert {
				newNodeArray = append(newNodeArray, convertedItem)
			}
		}

		// Check relationships array
		newRelationshipArray := make([]neo4j.Relationship, 0)
		for _, item := range typedVal {
			if convertedItem, canConvert := item.(neo4j.Relationship); canConvert {
				newRelationshipArray = append(newRelationshipArray, convertedItem)
			}
		}

		// Check paths array
		newPathArray := make([]neo4j.Path, 0)
		for _, item := range typedVal {
			if convertedItem, canConvert := item.(neo4j.Path); canConvert {
				newPathArray = append(newPathArray, convertedItem)
			}
		}

		// Check which array could be decoded
		switch len(typedVal) {
		case len(newRecordMapArray):
			resultRecord.MapArrays[key] = newRecordMapArray
		case len(newNodeArray):
			resultRecord.NodeArrays[key] = newNodeArray
		case len(newRelationshipArray):
			resultRecord.RelationshipArrays[key] = newRelationshipArray
		case len(newPathArray):
			resultRecord.PathArrays[key] = newPathArray
		default:
			resultRecord.Arrays[key] = typedVal
		}
	case map[string]interface{}:
		innerRecordMap, err := decodeMap(typedVal)
		if err != nil {
			return err
		}
		resultRecord.Maps[key] = *innerRecordMap
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

	return nil
}

// decodeMap takes a map as input and converts it into a RecordMap for simplicity of use
func decodeMap(mapInterface map[string]interface{}) (*RecordMap, internalErr.Neo4GoError) {
	newRecordMap := newEmptyRecordMap()

	for key, val := range mapInterface {
		if val == nil {
			continue
		}

		err := decodeItemInRecordMap(key, val, &newRecordMap)
		if err != nil {
			return nil, err
		}
	}

	return &newRecordMap, nil
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
func (res *queryResult) Record() (*RecordMap, internalErr.Neo4GoError) {
	record := res.result.Record()

	newRecordMap := newEmptyRecordMap()
	for _, recordKey := range record.Keys() {
		recordValue, exists := record.Get(recordKey)
		if !exists || recordValue == nil {
			continue
		}

		err := decodeItemInRecordMap(recordKey, recordValue, &newRecordMap)
		if err != nil {
			return nil, err
		}
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
