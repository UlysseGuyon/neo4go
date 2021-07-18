package neo4go

import (
	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
)

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
