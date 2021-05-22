package errors

import "fmt"

type Neo4GoQueryBuildError struct {
	Bare   bool
	Reason string
}

func (err Neo4GoQueryBuildError) Error() string {
	if err.Bare {
		return errorFmt("Query Building", fmt.Sprintf(
			"Could not build Neo4J query because %s",
			err.Reason,
		))
	} else {
		return errorFmt("Query Building", err.Reason)
	}
}

func (err Neo4GoQueryBuildError) IsConnError() bool       { return false }
func (err Neo4GoQueryBuildError) IsInitError() bool       { return false }
func (err Neo4GoQueryBuildError) IsQueryBuildError() bool { return true }
func (err Neo4GoQueryBuildError) IsQueryError() bool      { return false }
func (err Neo4GoQueryBuildError) IsUnknownError() bool    { return false }
