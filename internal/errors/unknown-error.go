package errors

type Neo4JUnknownError struct{}

func (err Neo4JUnknownError) Error() string {
	return errorFmt("Unknown", "An unknown error has occured")
}

func (err Neo4JUnknownError) IsConnError() bool       { return false }
func (err Neo4JUnknownError) IsInitError() bool       { return false }
func (err Neo4JUnknownError) IsQueryBuildError() bool { return false }
func (err Neo4JUnknownError) IsQueryError() bool      { return false }
func (err Neo4JUnknownError) IsUnknownError() bool    { return true }
