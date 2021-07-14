package neo4go

import (
	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
)

func IsInitError(err error) bool {
	_, canConvert := err.(*internalErr.InitError)
	return canConvert
}

func IsTypeError(err error) bool {
	_, canConvert := err.(*internalErr.TypeError)
	return canConvert
}

func IsDecodingError(err error) bool {
	_, canConvert := err.(*internalErr.DecodingError)
	return canConvert
}

func IsQueryError(err error) bool {
	_, canConvert := err.(*internalErr.QueryError)
	return canConvert
}

func IsUnknownError(err error) bool {
	_, canConvert := err.(*internalErr.UnknownError)
	return canConvert
}

func IsSecurityError(err error) bool {
	_, canConvert := err.(*internalErr.SecurityError)
	return canConvert
}

func IsAuthError(err error) bool {
	_, canConvert := err.(*internalErr.AuthError)
	return canConvert
}

func IsClientError(err error) bool {
	_, canConvert := err.(*internalErr.ClientError)
	return canConvert
}

func IsTransientError(err error) bool {
	_, canConvert := err.(*internalErr.TransientError)
	return canConvert
}

func IsSessionError(err error) bool {
	_, canConvert := err.(*internalErr.SessionError)
	return canConvert
}

func IsServiceUnavailableError(err error) bool {
	_, canConvert := err.(*internalErr.UnavailableError)
	return canConvert
}
