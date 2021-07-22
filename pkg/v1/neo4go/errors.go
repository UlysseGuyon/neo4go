package neo4go

import (
	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Neo4GoError is an error interface allowing to point out more precisely the error type
type Neo4GoError interface {
	error

	// FmtError returns the formatted error with a prefix
	FmtError() string
}

// ToDriverError converts a raw error to a Neo4GoError, searching first for neo4j-go-driver errors
func ToDriverError(err error) Neo4GoError {
	// First check for all the neo4j-go-driver errors
	if neo4j.IsSecurityError(err) {
		return &internalErr.SecurityError{
			Err: err.Error(),
		}
	} else if neo4j.IsAuthenticationError(err) {
		return &internalErr.AuthError{
			Err: err.Error(),
		}
	} else if neo4j.IsClientError(err) {
		return &internalErr.ClientError{
			Err: err.Error(),
		}
	} else if neo4j.IsTransientError(err) {
		return &internalErr.TransientError{
			Err: err.Error(),
		}
	} else if neo4j.IsSessionExpired(err) {
		return &internalErr.SessionError{
			Err: err.Error(),
		}
	} else if neo4j.IsServiceUnavailable(err) {
		return &internalErr.UnavailableError{
			Err: err.Error(),
		}
	} else if convertedErr, canConvert := err.(Neo4GoError); canConvert {
		// Then if the error is already a Neo4GoError, return it as it is
		return convertedErr
	}

	// Finally, return an unknown error if nothing was found
	return &internalErr.UnknownError{
		Err: err.Error(),
	}
}

// IsInitError tells if the error is a neo4go Init error
func IsInitError(err error) bool {
	_, canConvert := err.(*internalErr.InitError)
	return canConvert
}

// IsTypeError tells if the error is a neo4go Type error
func IsTypeError(err error) bool {
	_, canConvert := err.(*internalErr.TypeError)
	return canConvert
}

// IsDecodingError tells if the error is a neo4go Decoding error
func IsDecodingError(err error) bool {
	_, canConvert := err.(*internalErr.DecodingError)
	return canConvert
}

// IsQueryError tells if the error is a neo4go Query error
func IsQueryError(err error) bool {
	_, canConvert := err.(*internalErr.QueryError)
	return canConvert
}

// IsTransactionError tells if the error is a neo4go Query error
func IsTransactionError(err error) bool {
	_, canConvert := err.(*internalErr.TransactionError)
	return canConvert
}

// IsUnknownError tells if the error is a neo4go Unknown error
func IsUnknownError(err error) bool {
	_, canConvert := err.(*internalErr.UnknownError)
	return canConvert
}

// IsSecurityError tells if the error is a neo4go Security error
func IsSecurityError(err error) bool {
	_, canConvert := err.(*internalErr.SecurityError)
	return canConvert
}

// IsAuthError tells if the error is a neo4go Auth error
func IsAuthError(err error) bool {
	_, canConvert := err.(*internalErr.AuthError)
	return canConvert
}

// IsClientError tells if the error is a neo4go Client error
func IsClientError(err error) bool {
	_, canConvert := err.(*internalErr.ClientError)
	return canConvert
}

// IsTransientError tells if the error is a neo4go Transient error
func IsTransientError(err error) bool {
	_, canConvert := err.(*internalErr.TransientError)
	return canConvert
}

// IsSessionError tells if the error is a neo4go Session error
func IsSessionError(err error) bool {
	_, canConvert := err.(*internalErr.SessionError)
	return canConvert
}

// IsServiceUnavailableError tells if the error is a neo4go Service Unavailable error
func IsServiceUnavailableError(err error) bool {
	_, canConvert := err.(*internalErr.UnavailableError)
	return canConvert
}
