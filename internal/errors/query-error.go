package errors

import "fmt"

type Neo4JQueryError struct {
	Bare   bool
	Reason string
}

func (err Neo4JQueryError) Error() string {
	if err.Bare {
		return errorFmt("Query", fmt.Sprintf(
			"The Neo4J query could not be run because %s",
			err.Reason,
		))
	} else {
		return errorFmt("Query", err.Reason)
	}
}

func (err Neo4JQueryError) IsConnError() bool       { return false }
func (err Neo4JQueryError) IsInitError() bool       { return false }
func (err Neo4JQueryError) IsQueryBuildError() bool { return false }
func (err Neo4JQueryError) IsQueryError() bool      { return true }
func (err Neo4JQueryError) IsUnknownError() bool    { return false }
