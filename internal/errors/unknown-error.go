package errors

type Neo4GoUnknownError struct{}

func (err Neo4GoUnknownError) Error() string {
	return errorFmt("Unknown", "An unknown error has occured")
}

func (err Neo4GoUnknownError) IsConnError() bool       { return false }
func (err Neo4GoUnknownError) IsInitError() bool       { return false }
func (err Neo4GoUnknownError) IsQueryBuildError() bool { return false }
func (err Neo4GoUnknownError) IsQueryError() bool      { return false }
func (err Neo4GoUnknownError) IsUnknownError() bool    { return true }
