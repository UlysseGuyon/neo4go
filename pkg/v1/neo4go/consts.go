package neo4go

import "github.com/UlysseGuyon/neo4go/internal/neo4go"

const (
	ConcurrencyModeMultiWrite        = neo4go.ConcurrencyModeMultiWrite
	ConcurrencyModeOneWriteMultiRead = neo4go.ConcurrencyModeOneWriteMultiRead
	ConcurrencyModeOnlyMultiRead     = neo4go.ConcurrencyModeOnlyMultiRead
	ConcurrencyModeNoConcurrency     = neo4go.ConcurrencyModeNoConcurrency
)
