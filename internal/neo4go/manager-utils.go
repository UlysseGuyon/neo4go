package neo4go

import (
	"strings"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	internalTypes "github.com/UlysseGuyon/neo4go/internal/types"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func CloseSessionChanel(ch chan neo4j.Session) error {
	close(ch)

	for session, ok := <-ch; ok; session, ok = <-ch {
		err := session.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func ValidateManagerOptions(opt internalTypes.ManagerOptions) internalErr.Neo4GoError {
	if opt.URI == "" {
		return internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: "Database URI given in options is empty",
		}
	}

	if opt.DatabaseName == "" {
		return internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: "Database name given in options is empty",
		}
	}

	return nil
}

func SetManagerOptionsDefaultValues(opt internalTypes.ManagerOptions) internalTypes.ManagerOptions {
	if opt.Concurrency == 0 {
		opt.Concurrency = ConcurrencyModeOneWriteMultiRead
	}

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
