package neo4go

const (
	ConcurrencyModeMultiWrite uint = 1 + iota
	ConcurrencyModeOneWriteMultiRead
	ConcurrencyModeOnlyMultiRead
	ConcurrencyModeNoConcurrency
)

const (
	DefaultDecodingTagName = "neo4j"
)
