package neo4go

import (
	"strings"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	internalTypes "github.com/UlysseGuyon/neo4go/internal/types"
)

func ValidateManagerOptions(opt internalTypes.ManagerOptions) internalErr.Neo4GoError {
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

func SetManagerOptionsDefaultValues(opt internalTypes.ManagerOptions) internalTypes.ManagerOptions {
	return opt
}

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
