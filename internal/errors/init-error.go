package errors

import "fmt"

type Neo4GoInitError struct {
	Bare   bool
	Reason string
}

func (err Neo4GoInitError) Error() string {
	if err.Bare {
		return errorFmt("Init", fmt.Sprintf(
			"Could not initialize Neo4J Database Manager because %s",
			err.Reason,
		))
	} else {
		return errorFmt("Init", err.Reason)
	}
}

func (err Neo4GoInitError) IsConnError() bool       { return false }
func (err Neo4GoInitError) IsInitError() bool       { return true }
func (err Neo4GoInitError) IsQueryBuildError() bool { return false }
func (err Neo4GoInitError) IsQueryError() bool      { return false }
func (err Neo4GoInitError) IsUnknownError() bool    { return false }
