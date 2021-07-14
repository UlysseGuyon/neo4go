package errors

type Neo4GoError interface {
	error
	FmtError() string
}

type InitError struct {
	Err string
}

func (err *InitError) Error() string {
	return err.Err
}

func (err *InitError) FmtError() string {
	return errorFmt("Init", err.Err)
}

type TypeError struct {
	Err string
}

func (err *TypeError) Error() string {
	return err.Err
}

func (err *TypeError) FmtError() string {
	return errorFmt("Type", err.Err)
}

type DecodingError struct {
	Err string
}

func (err *DecodingError) Error() string {
	return err.Err
}

func (err *DecodingError) FmtError() string {
	return errorFmt("Decoding", err.Err)
}

type QueryError struct {
	Err string
}

func (err *QueryError) Error() string {
	return err.Err
}

func (err *QueryError) FmtError() string {
	return errorFmt("Query", err.Err)
}

type UnknownError struct {
	Err string
}

func (err *UnknownError) Error() string {
	return err.Err
}

func (err *UnknownError) FmtError() string {
	return errorFmt("Unknown", err.Err)
}
