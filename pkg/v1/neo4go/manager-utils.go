package neo4go

import (
	"strings"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
)

// validateManagerOptions allows early detection of wrong options
func validateManagerOptions(opt ManagerOptions) Neo4GoError {
	if opt.URI == "" {
		return &internalErr.InitError{
			Err:    "Database URI given in options is empty",
			DBName: opt.DatabaseName,
			URI:    opt.URI,
		}
	}

	if opt.DatabaseName == "" {
		return &internalErr.InitError{
			Err:    "Database name given in options is empty",
			DBName: opt.DatabaseName,
			URI:    opt.URI,
		}
	}

	return nil
}

// IsWriteQuery determines if the given cypher query's type is write or read
func IsWriteQuery(query string) bool {
	isWrite := false

	allWriteClauses := []string{
		"CREATE",
		"MERGE",
		"DELETE",
		"SET",
		"REMOVE",
		"FOREACH",
		"DROP",
		"ALTER",
		"RENAME",
		"GRANT",
		"REVOKE",
		"DENY",
	}

	for _, writeClause := range allWriteClauses {
		if strings.Contains(strings.ToLower(query), strings.ToLower(writeClause)) {
			isWrite = true
		}
	}

	return isWrite
}
