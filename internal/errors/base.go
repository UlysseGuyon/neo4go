package errors

type Neo4GoError interface {
	error
	IsConnError() bool
	IsInitError() bool
	IsQueryBuildError() bool
	IsQueryError() bool
	IsUnknownError() bool
}
