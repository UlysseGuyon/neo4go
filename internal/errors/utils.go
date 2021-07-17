package errors

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// errorFmt formats an error string by showing its type as prefix
func errorFmt(t string, err string) string {
	return fmt.Sprintf(
		`[Neo4Go %s Error]: %s`,
		t,
		err,
	)
}

// ToDriverError converts a raw error to a Neo4GoError, searching first for neo4j-go-driver errors
func ToDriverError(err error) Neo4GoError {
	// First check for all the neo4j-go-driver errors
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
	} else if convertedErr, canConvert := err.(Neo4GoError); canConvert {
		// Then if the error is already a Neo4GoError, return it as it is
		return convertedErr
	}

	// Finally, return an unknown error if nothing was found
	return &UnknownError{
		Err: err.Error(),
	}
}
