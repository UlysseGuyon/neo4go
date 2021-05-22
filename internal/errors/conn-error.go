package errors

import "fmt"

type Neo4GoConnError struct {
	URI    string
	DBName string
}

func (err Neo4GoConnError) Error() string {
	return errorFmt("Connexion", fmt.Sprintf(
		"Could not connect to database `%s/%s`",
		err.URI,
		err.DBName,
	))
}

func (err Neo4GoConnError) IsConnError() bool       { return true }
func (err Neo4GoConnError) IsInitError() bool       { return false }
func (err Neo4GoConnError) IsQueryBuildError() bool { return false }
func (err Neo4GoConnError) IsQueryError() bool      { return false }
func (err Neo4GoConnError) IsUnknownError() bool    { return false }
