package neo4go

import (
	"time"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type QueryParams struct {
	Query       string
	Params      map[string]InputObject
	Configurers []func(*neo4j.TransactionConfig)
	Bookmarks   []string
}

type TransactionStepParams struct {
	Query          string
	Params         map[string]InputObject
	TransitionFunc func(QueryResult) map[string]InputObject
}

type TransactionParams struct {
	TransactionSteps []TransactionStepParams
	Configurers      []func(*neo4j.TransactionConfig)
	Bookmarks        []string
}

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

type Manager interface {
	Init(ManagerOptions) internalErr.Neo4GoError
	IsConnected() bool
	Close() internalErr.Neo4GoError
	Query(QueryParams) (QueryResult, internalErr.Neo4GoError)
	Transaction(TransactionParams) (QueryResult, internalErr.Neo4GoError)
}

type InputObject interface {
	Convert() map[string]InputObject
}
