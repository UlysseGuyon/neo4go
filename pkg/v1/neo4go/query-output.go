package neo4go

import (
	"time"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type RecordMap struct {
	Arrays    map[string][]interface{}
	MapArrays map[string][]RecordMap
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

func newEmptyRecordMap() RecordMap {
	return RecordMap{
		Arrays:    make(map[string][]interface{}),
		MapArrays: make(map[string][]RecordMap),
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

type QueryResult interface {
	// Keys returns the keys available on the result set.
	Keys() ([]string, error)
	// Next returns true only if there is a record to be processed.
	Next() bool
	// Err returns the latest error that caused this Next to return false.
	Err() error
	// Record returns the current record.
	Record() (*RecordMap, internalErr.Neo4GoError)
	// Summary returns the summary information about the statement execution.
	Summary() (neo4j.ResultSummary, error)
	// Consume consumes the entire result and returns the summary information
	// about the statement execution.
	Consume() (neo4j.ResultSummary, error)

	RawResult() neo4j.Result
}

// Single returns one and only one record from the result stream. Any error passed in
// or reported while navigating the result stream is returned without any conversion.
// If the result stream contains zero or more than one records error is returned.
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

func decodeItemInRecordMap(key string, value interface{}, resultRecord *RecordMap) internalErr.Neo4GoError {
	switch typedVal := value.(type) {
	case []interface{}:
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

		if len(newRecordMapArray) == len(typedVal) {
			resultRecord.MapArrays[key] = newRecordMapArray
		} else {
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

type queryResult struct {
	Result neo4j.Result
}

func newQueryResult(result neo4j.Result) QueryResult {
	return &queryResult{Result: result}
}

func (res *queryResult) Keys() ([]string, error) {
	return res.Result.Keys()
}
func (res *queryResult) Next() bool {
	return res.Result.Next()
}
func (res *queryResult) Err() error {
	return res.Result.Err()
}
func (res *queryResult) Record() (*RecordMap, internalErr.Neo4GoError) {
	record := res.Result.Record()

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
func (res *queryResult) Summary() (neo4j.ResultSummary, error) {
	return res.Result.Summary()
}
func (res *queryResult) Consume() (neo4j.ResultSummary, error) {
	return res.Result.Consume()
}
func (res *queryResult) RawResult() neo4j.Result {
	return res.Result
}
