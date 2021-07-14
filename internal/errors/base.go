package errors

import "fmt"

type Neo4GoError interface {
	error
	FmtError() string
}

type InitError struct {
	Err    string
	URI    string
	DBName string
}

func (err *InitError) Error() string {
	return fmt.Sprintf("%s (URI : %s / Database : %s)", err.Err, err.URI, err.DBName)
}

func (err *InitError) FmtError() string {
	return errorFmt("Init", err.Error())
}

type TypeError struct {
	Err           string
	ExpectedTypes []string
	GotType       string
}

func (err *TypeError) Error() string {
	return fmt.Sprintf("%s (Expected types : %v / Got : %s)", err.Err, err.ExpectedTypes, err.GotType)
}

func (err *TypeError) FmtError() string {
	return errorFmt("Type", err.Error())
}

type DecodingError struct {
	Err string
}

func (err *DecodingError) Error() string {
	return err.Err
}

func (err *DecodingError) FmtError() string {
	return errorFmt("Decoding", err.Error())
}

type QueryError struct {
	Err string
}

func (err *QueryError) Error() string {
	return err.Err
}

func (err *QueryError) FmtError() string {
	return errorFmt("Query", err.Error())
}

type UnknownError struct {
	Err string
}

func (err *UnknownError) Error() string {
	return err.Err
}

func (err *UnknownError) FmtError() string {
	return errorFmt("Unknown", err.Error())
}
