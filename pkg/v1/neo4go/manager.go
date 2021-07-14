package neo4go

import (
	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	internalMain "github.com/UlysseGuyon/neo4go/internal/neo4go"
	internalTypes "github.com/UlysseGuyon/neo4go/internal/types"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type manager struct {
	options *internalTypes.ManagerOptions
	driver  *neo4j.Driver
}

func NewManager(options ManagerOptions) (Manager, internalErr.Neo4GoError) {
	newManager := manager{}

	err := newManager.Init(options)
	if err != nil {
		return nil, err
	}

	if !newManager.IsConnected() {
		return nil, internalErr.Neo4GoConnError{
			URI:    options.URI,
			DBName: options.DatabaseName,
		}
	}

	return &newManager, nil
}

func (m *manager) Init(options ManagerOptions) internalErr.Neo4GoError {
	optErr := internalMain.ValidateManagerOptions(internalTypes.ManagerOptions(options))
	if optErr != nil {
		return optErr
	}
	usedOptions := internalMain.SetManagerOptionsDefaultValues(internalTypes.ManagerOptions(options))

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
		return internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: err.Error(),
		}
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
		return internalErr.Neo4GoUnknownError{}
	}

	return nil
}

func (m *manager) Query(queryParams QueryParams) (QueryResult, internalErr.Neo4GoError) {
	paramsMap := make(map[string]interface{})
	for key, value := range queryParams.Params {
		paramsMap[key] = convertInputObject(value)
	}

	isWrite := internalMain.IsWriteQuery(queryParams.Query)

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
		return nil, &internalErr.Neo4GoQueryError{
			Reason: err.Error(),
		}
	}
	defer session.Close()

	rawResult, err := session.Run(queryParams.Query, paramsMap, queryParams.Configurers...)
	if err != nil {
		return nil, &internalErr.Neo4GoQueryError{
			Reason: err.Error(),
		}
	}

	convertedResult := newQueryResult(rawResult)

	return convertedResult, nil
}

func (m *manager) Transaction(transactionGlobalParams TransactionParams) (QueryResult, internalErr.Neo4GoError) {
	transactionWork := func(tx neo4j.Transaction) (interface{}, error) {
		var nextQueryParams map[string]InputObject = nil
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
		if internalMain.IsWriteQuery(transactionParams.Query) {
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
		return nil, &internalErr.Neo4GoQueryError{
			Reason: err.Error(),
		}
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
		return nil, &internalErr.Neo4GoQueryError{
			Reason: transactionErr.Error(),
		}
	}

	transactionResult, canConvert := transactionResultI.(QueryResult)
	if !canConvert {
		return nil, &internalErr.Neo4GoQueryError{
			Reason: "Could not convert transaction result to structured QueryResult",
		}
	}

	return transactionResult, nil
}
