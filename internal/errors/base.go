package errors

import "fmt"

// Neo4GoError is an error interface allowing to point out more precisely the error type
type Neo4GoError interface {
	error

	// FmtError returns the formatted error with a prefix
	FmtError() string
}

// The strings representing each type of neo4go native error
const (
	InitErrorTypeName     = "Init"
	TypeErrorTypeName     = "Type"
	DecodingErrorTypeName = "Decoding"
	QueryErrorTypeName    = "Query"
	UnknownErrorTypeName  = "Unknown"
)

/* ----- INIT ERROR ----- */

// InitError represents an error occurring at the init of the manager
type InitError struct {
	// The basic error string
	Err string

	// The URI used in the manager
	URI string

	// The database name used in the manager
	DBName string
}

// Error returns the raw error string
func (err *InitError) Error() string {
	return fmt.Sprintf("%s (URI : %s / Database : %s)", err.Err, err.URI, err.DBName)
}

// FmtError returns the formatted error string
func (err *InitError) FmtError() string {
	return errorFmt(InitErrorTypeName, err.Error())
}

/* ----- TYPE ERROR ----- */

// TypeError represents an error occurring on a type convertion
type TypeError struct {
	// The basic error string
	Err string

	// The expected types that would not have triggered the error
	ExpectedTypes []string

	// The type that triggered the error
	GotType string
}

// Error returns the raw error string
func (err *TypeError) Error() string {
	return fmt.Sprintf("%s (Expected types : %v / Got : %s)", err.Err, err.ExpectedTypes, err.GotType)
}

// FmtError returns the formatted error string
func (err *TypeError) FmtError() string {
	return errorFmt("Type", err.Error())
}

/* ----- DECODING ERROR ----- */

// DecodingError represents an error occurring durring the decoding of a Neo4J result into a custom user struct
type DecodingError struct {
	Err string
}

// Error returns the raw error string
func (err *DecodingError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *DecodingError) FmtError() string {
	return errorFmt("Decoding", err.Error())
}

/* ----- QUERY ERROR ----- */

// QueryError represents an error occurring durring the execution of a Neo4J query
type QueryError struct {
	Err string
}

// Error returns the raw error string
func (err *QueryError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *QueryError) FmtError() string {
	return errorFmt("Query", err.Error())
}

/* ----- UNKOWN ERROR ----- */

// UnknownError represents any error not known by the neo4go package
type UnknownError struct {
	Err string
}

// Error returns the raw error string
func (err *UnknownError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *UnknownError) FmtError() string {
	return errorFmt("Unknown", err.Error())
}
