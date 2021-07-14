package errors

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func errorFmt(t string, err string) string {
	return fmt.Sprintf(
		`[Neo4Go %s Error]: %s`,
		t,
		err,
	)
}

func ToDriverError(err error) Neo4GoError {
	if neo4j.IsSecurityError(err) {
		return &SecurityError{
			Err: err.Error(),
		}
	} else if neo4j.IsAuthenticationError(err) {
		return &AuthError{
			Err: err.Error(),
		}
	} else if neo4j.IsClientError(err) {
		return &ClientError{
			Err: err.Error(),
		}
	} else if neo4j.IsTransientError(err) {
		return &TransientError{
			Err: err.Error(),
		}
	} else if neo4j.IsSessionExpired(err) {
		return &SessionError{
			Err: err.Error(),
		}
	} else if neo4j.IsServiceUnavailable(err) {
		return &UnavailableError{
			Err: err.Error(),
		}
	}

	return &UnknownError{
		Err: err.Error(),
	}
}
