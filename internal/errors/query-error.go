package errors

import "fmt"

type Neo4GoQueryError struct {
	Bare   bool
	Reason string
}

func (err Neo4GoQueryError) Error() string {
	if err.Bare {
		return errorFmt("Query", fmt.Sprintf(
			"The Neo4J query could not be run because %s",
			err.Reason,
		))
	} else {
		return errorFmt("Query", err.Reason)
	}
}

func (err Neo4GoQueryError) IsConnError() bool       { return false }
func (err Neo4GoQueryError) IsInitError() bool       { return false }
func (err Neo4GoQueryError) IsQueryBuildError() bool { return false }
func (err Neo4GoQueryError) IsQueryError() bool      { return true }
func (err Neo4GoQueryError) IsUnknownError() bool    { return false }
