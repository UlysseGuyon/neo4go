package neo4go

import (
	"fmt"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Manager is a wrapper around the neo4j-go-driver that simplifies its usage and adds type checking
type Manager interface {
	// IsConnected tells if the driver could effectively connect to the database
	IsConnected() bool

	// Query allows a single query transaction to be made in database
	Query(QueryParams) (QueryResult, internalErr.Neo4GoError)

	// Transaction allows multiple queries to be made in database inside a single transaction that rools back in case of error
	Transaction(TransactionParams) (QueryResult, internalErr.Neo4GoError)

	// Close closes the driver
	Close() internalErr.Neo4GoError
}

// ManagerOptions represents the configuration applied to a manager
type ManagerOptions struct {
	// The full URI to access the database (of the form protocol://domain:port)
	URI string

	// The name of the database to use
	DatabaseName string

	// The realm to use
	Realm string

	// The username to use for accessing the database with basic auth
	Username string

	// The password to use for accessing the database with basic auth
	Password string

	// The auth token used for any other kind of authentication than basic auth
	CustomAuth *neo4j.AuthToken

	// The neo4j configurers to apply to the driver
	Configurers []func(*neo4j.Config)
}

// QueryParams represents all the configuration of a single query transaction
type QueryParams struct {
	// The cypher query to run in database
	Query string

	// The parameters to apply to this query
	Params map[string]InputStruct

	// The configurers to apply to the query
	Configurers []func(*neo4j.TransactionConfig)

	// The bookmarks of previous sessions to apply to the query
	Bookmarks []string
}

// TransactionStepParams represents all the configuration of a single query inside a multiple queries transaction
type TransactionStepParams struct {
	// The cypher query to run in database
	Query string

	// The parameters to apply to this query
	Params map[string]InputStruct

	// The function to apply on the result of this query to obtain the parameters of the next one
	TransitionFunc func(QueryResult) (map[string]InputStruct, error)
}

// TransactionParams represents all the configuration of a whole multiple queries transaction
type TransactionParams struct {
	// The list of individual queries forming this transaction
	TransactionSteps []TransactionStepParams

	// The configurers to apply to this transaction
	Configurers []func(*neo4j.TransactionConfig)

	// The bookmarks of previous sessions to apply to the transaction
	Bookmarks []string
}

// manager is the default implementation of the Manager interface
type manager struct {
	// The configuration of this manager
	options *ManagerOptions

	// The neo4j-go-driver wrapped in this manager
	driver *neo4j.Driver
}

// NewManager creates a new instance of Manager, with a given config.
func NewManager(options ManagerOptions) (Manager, internalErr.Neo4GoError) {
	newManager := manager{}

	// Init the driver of the manager
	err := newManager.init(options)
	if err != nil {
		return nil, err
	}

	// Test the connection of the driver
	if !newManager.IsConnected() {
		return nil, &internalErr.InitError{
			Err:    "Could not connect to database",
			URI:    options.URI,
			DBName: options.DatabaseName,
		}
	}

	return &newManager, nil
}

// init creates a new driver in the manager from the given config.
func (m *manager) init(options ManagerOptions) internalErr.Neo4GoError {
	// Check if the config was correctly filled
	optErr := validateManagerOptions(options)
	if optErr != nil {
		return optErr
	}

	// Select the right auth protocol
	usedAuth := neo4j.NoAuth()
	if options.CustomAuth != nil {
		usedAuth = *options.CustomAuth
	} else if options.Username != "" && options.Password != "" {
		usedAuth = neo4j.BasicAuth(options.Username, options.Password, options.Realm)
	}

	// Create the driver
	newDriver, err := neo4j.NewDriver(
		options.URI,
		usedAuth,
		options.Configurers...,
	)

	if err != nil {
		return internalErr.ToDriverError(err)
	}

	m.options = &options
	m.driver = &newDriver

	return nil
}

// IsConnected tells if the driver could effectively connect to the database
func (m *manager) IsConnected() bool {
	if m.driver == nil {
		return false
	}

	err := (*m.driver).VerifyConnectivity()

	return err == nil
}

// Close closes the driver
func (m *manager) Close() internalErr.Neo4GoError {
	err := (*m.driver).Close()
	if err != nil {
		return internalErr.ToDriverError(err)
	}

	return nil
}

// Query allows a single query transaction to be made in database
func (m *manager) Query(queryParams QueryParams) (QueryResult, internalErr.Neo4GoError) {
	// First, we convert all the input objects as interface maps
	paramsMap := make(map[string]interface{})
	for key, value := range queryParams.Params {
		paramsMap[key] = convertInputObject(value)
	}

	// Determine if the query is read or write and set the access mode depending on it
	isWrite := isWriteQuery(queryParams.Query)

	usedSessionMode := neo4j.AccessModeRead
	if isWrite {
		usedSessionMode = neo4j.AccessModeWrite
	}

	// Create the new session from configuration
	session, err := (*m.driver).NewSession(neo4j.SessionConfig{
		AccessMode:   usedSessionMode,
		DatabaseName: m.options.DatabaseName,
		Bookmarks:    queryParams.Bookmarks,
	})
	if err != nil {
		return nil, internalErr.ToDriverError(err)
	}
	defer session.Close()

	// Run the query with the new session and the query config
	rawResult, err := session.Run(queryParams.Query, paramsMap, queryParams.Configurers...)
	if err != nil {
		return nil, internalErr.ToDriverError(err)
	}

	// Convert the raw results as a collection of typed results
	convertedResult := newQueryResult(rawResult)

	return convertedResult, nil
}

// Transaction allows multiple queries to be made in database inside a single transaction that rools back in case of error
func (m *manager) Transaction(transactionGlobalParams TransactionParams) (QueryResult, internalErr.Neo4GoError) {
	// Create the transaction work
	transactionWork := func(tx neo4j.Transaction) (interface{}, error) {
		var nextQueryParams map[string]InputStruct = nil
		var lastResult QueryResult

		// Loop through all the queries of the transaction
		for _, transactionParams := range transactionGlobalParams.TransactionSteps {
			// If there was a previous query that set the params of this one, then take them. Else take the ones in the config
			usedParamsMap := transactionParams.Params
			if nextQueryParams != nil {
				usedParamsMap = nextQueryParams
			}

			// Convert all the input objects as interface maps
			paramsMap := make(map[string]interface{})
			for key, value := range usedParamsMap {
				paramsMap[key] = convertInputObject(value)
			}

			// Run the individual query
			result, err := tx.Run(transactionParams.Query, paramsMap)
			if err != nil {
				rollErr := tx.Rollback()
				if rollErr != nil {
					return nil, rollErr
				}
				return nil, err
			}

			// Convert the raw results as a collection of typed results
			lastResult = newQueryResult(result)

			// If there is a transition function to determine the next query params, then run it and save those params
			if transactionParams.TransitionFunc != nil {
				nextQueryParams, err = transactionParams.TransitionFunc(lastResult)
				if err != nil {
					rollErr := tx.Rollback()
					if rollErr != nil {
						return nil, rollErr
					}
					return nil, err
				}
			}
		}

		// In the end, commit the transaction and return the last result
		err := tx.Commit()
		if err != nil {
			return nil, err
		}

		return lastResult, nil
	}

	// If one of the queries is a write, then we consider the whole transaction as a write
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

	// Create the new session from configuration
	session, err := (*m.driver).NewSession(neo4j.SessionConfig{
		AccessMode:   usedSessionMode,
		DatabaseName: m.options.DatabaseName,
		Bookmarks:    transactionGlobalParams.Bookmarks,
	})
	if err != nil {
		return nil, internalErr.ToDriverError(err)
	}
	defer session.Close()

	// Run the transaction
	var transactionResultI interface{}
	var transactionErr error
	if isWrite {
		transactionResultI, transactionErr = session.WriteTransaction(transactionWork, transactionGlobalParams.Configurers...)
	} else {
		transactionResultI, transactionErr = session.ReadTransaction(transactionWork, transactionGlobalParams.Configurers...)
	}

	if transactionErr != nil {
		return nil, internalErr.ToDriverError(transactionErr)
	}

	// Cast the already converted typed result
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
