package neo4go

import (
	"fmt"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Manager interface {
	Init(ManagerOptions) internalErr.Neo4GoError
	IsConnected() bool
	Close() internalErr.Neo4GoError
	Query(QueryParams) (QueryResult, internalErr.Neo4GoError)
	Transaction(TransactionParams) (QueryResult, internalErr.Neo4GoError)
}

type ManagerOptions struct {
	URI          string
	DatabaseName string
	Realm        string
	Username     string
	Password     string
	CustomAuth   *neo4j.AuthToken
	Configurers  []func(*neo4j.Config)
}

type QueryParams struct {
	Query       string
	Params      map[string]InputStruct
	Configurers []func(*neo4j.TransactionConfig)
	Bookmarks   []string
}

type TransactionStepParams struct {
	Query          string
	Params         map[string]InputStruct
	TransitionFunc func(QueryResult) map[string]InputStruct
}

type TransactionParams struct {
	TransactionSteps []TransactionStepParams
	Configurers      []func(*neo4j.TransactionConfig)
	Bookmarks        []string
}

type manager struct {
	options *ManagerOptions
	driver  *neo4j.Driver
}

func NewManager(options ManagerOptions) (Manager, internalErr.Neo4GoError) {
	newManager := manager{}

	err := newManager.Init(options)
	if err != nil {
		return nil, err
	}

	if !newManager.IsConnected() {
		return nil, &internalErr.InitError{
			Err:    "Could not connect to database",
			URI:    options.URI,
			DBName: options.DatabaseName,
		}
	}

	return &newManager, nil
}

func (m *manager) Init(options ManagerOptions) internalErr.Neo4GoError {
	optErr := validateManagerOptions(options)
	if optErr != nil {
		return optErr
	}
	usedOptions := setManagerOptionsDefaultValues(options)

	usedAuth := neo4j.NoAuth()
	if usedOptions.CustomAuth != nil {
		usedAuth = *usedOptions.CustomAuth
	} else if usedOptions.Username != "" && usedOptions.Password != "" {
		usedAuth = neo4j.BasicAuth(usedOptions.Username, usedOptions.Password, usedOptions.Realm)
	}

	newDriver, err := neo4j.NewDriver(
		usedOptions.URI,
		usedAuth,
		usedOptions.Configurers...,
	)

	if err != nil {
		return internalErr.ToDriverError(err)
	}

	m.options = &usedOptions
	m.driver = &newDriver

	return nil
}

func (m *manager) IsConnected() bool {
	if m.driver == nil {
		return false
	}

	err := (*m.driver).VerifyConnectivity()

	return err == nil
}

func (m *manager) Close() internalErr.Neo4GoError {
	err := (*m.driver).Close()
	if err != nil {
		return internalErr.ToDriverError(err)
	}

	return nil
}

func (m *manager) Query(queryParams QueryParams) (QueryResult, internalErr.Neo4GoError) {
	paramsMap := make(map[string]interface{})
	for key, value := range queryParams.Params {
		paramsMap[key] = convertInputObject(value)
	}

	isWrite := isWriteQuery(queryParams.Query)

	usedSessionMode := neo4j.AccessModeRead
	if isWrite {
		usedSessionMode = neo4j.AccessModeWrite
	}
	session, err := (*m.driver).NewSession(neo4j.SessionConfig{
		AccessMode:   usedSessionMode,
		DatabaseName: m.options.DatabaseName,
		Bookmarks:    queryParams.Bookmarks,
	})
	if err != nil {
		return nil, internalErr.ToDriverError(err)
	}
	defer session.Close()

	rawResult, err := session.Run(queryParams.Query, paramsMap, queryParams.Configurers...)
	if err != nil {
		return nil, internalErr.ToDriverError(err)
	}

	convertedResult := newQueryResult(rawResult)

	return convertedResult, nil
}

func (m *manager) Transaction(transactionGlobalParams TransactionParams) (QueryResult, internalErr.Neo4GoError) {
	transactionWork := func(tx neo4j.Transaction) (interface{}, error) {
		var nextQueryParams map[string]InputStruct = nil
		var lastResult QueryResult
		for _, transactionParams := range transactionGlobalParams.TransactionSteps {
			usedParamsMap := transactionParams.Params
			if nextQueryParams != nil {
				usedParamsMap = nextQueryParams
			}

			paramsMap := make(map[string]interface{})
			for key, value := range usedParamsMap {
				paramsMap[key] = convertInputObject(value)
			}

			result, err := tx.Run(transactionParams.Query, paramsMap)
			if err != nil {
				rollErr := tx.Rollback()
				if rollErr != nil {
					return nil, rollErr
				}
				return nil, err
			}

			lastResult = newQueryResult(result)

			if transactionParams.TransitionFunc != nil {
				nextQueryParams = transactionParams.TransitionFunc(lastResult)
			}
		}

		err := tx.Commit()
		if err != nil {
			return nil, err
		}

		return lastResult, nil
	}

	isWrite := false
	for _, transactionParams := range transactionGlobalParams.TransactionSteps {
		if isWriteQuery(transactionParams.Query) {
			isWrite = true
		}
	}

	usedSessionMode := neo4j.AccessModeRead
	if isWrite {
		usedSessionMode = neo4j.AccessModeWrite
	}
	session, err := (*m.driver).NewSession(neo4j.SessionConfig{
		AccessMode:   usedSessionMode,
		DatabaseName: m.options.DatabaseName,
		Bookmarks:    transactionGlobalParams.Bookmarks,
	})
	if err != nil {
		return nil, internalErr.ToDriverError(err)
	}
	defer session.Close()

	var transactionResultI interface{}
	var transactionErr error
	if isWrite {
		transactionResultI, transactionErr = session.WriteTransaction(transactionWork, transactionGlobalParams.Configurers...)
	} else {
		transactionResultI, transactionErr = session.ReadTransaction(transactionWork, transactionGlobalParams.Configurers...)
	}

	if transactionErr != nil {
		return nil, internalErr.ToDriverError(err)
	}

	transactionResult, canConvert := transactionResultI.(QueryResult)
	if !canConvert {
		return nil, &internalErr.TypeError{
			Err:           "Could not convert transaction result to structured QueryResult",
			ExpectedTypes: []string{"QueryResult"},
			GotType:       fmt.Sprintf("%T", transactionResultI),
		}
	}

	return transactionResult, nil
}
