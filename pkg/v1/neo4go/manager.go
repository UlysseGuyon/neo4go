package neo4go

import (
	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	uuid "github.com/satori/go.uuid"
)

// Manager is a wrapper around the neo4j-go-driver that simplifies its usage and adds type checking
type Manager interface {
	// IsConnected tells if the driver could effectively connect to the database
	IsConnected() bool

	// Query allows a single query transaction to be made in database
	Query(QueryParams) (QueryResult, internalErr.Neo4GoError)

	BeginTransaction(TransactionParams) (string, internalErr.Neo4GoError)
	Commit(string) internalErr.Neo4GoError
	Rollback(string) internalErr.Neo4GoError

	LastBookmark() string

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

	Transaction string

	CommitOnSuccess bool
}

type TransactionParams struct {
	IsWrite bool

	Bookmarks []string

	Configurers []func(*neo4j.TransactionConfig)
}

type transactionSession struct {
	transaction neo4j.Transaction

	session neo4j.Session
}

// manager is the default implementation of the Manager interface
type manager struct {
	// The configuration of this manager
	options *ManagerOptions

	// The neo4j-go-driver wrapped in this manager
	driver *neo4j.Driver

	lastBookmark string

	transactionSessions map[string]transactionSession
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

	m.transactionSessions = make(map[string]transactionSession)

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
	for _, val := range m.transactionSessions {
		err := val.transaction.Close()
		if err != nil {
			return internalErr.ToDriverError(err)
		}

		err = val.session.Close()
		if err != nil {
			return internalErr.ToDriverError(err)
		}
	}

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

	txSession, useTransaction := m.transactionSessions[queryParams.Transaction]
	if !useTransaction && queryParams.Transaction != "" {
		return nil, &internalErr.TransactionError{
			Err: "Trying to query with a non existing transaction",
		}
	}

	var rawResult neo4j.Result
	var err error
	var usedSession neo4j.Session
	if useTransaction {
		usedSession = txSession.session

		rawResult, err = txSession.transaction.Run(queryParams.Query, paramsMap)
		if err != nil {
			return nil, internalErr.ToDriverError(err)
		}

		if queryParams.CommitOnSuccess {
			m.Commit(queryParams.Transaction)
		}
	} else {
		// Determine if the query is read or write and set the access mode depending on it
		isWrite := IsWriteQuery(queryParams.Query)

		usedSessionMode := neo4j.AccessModeRead
		if isWrite {
			usedSessionMode = neo4j.AccessModeWrite
		}

		// Create the new session from configuration
		usedSession, err = (*m.driver).NewSession(neo4j.SessionConfig{
			AccessMode:   usedSessionMode,
			DatabaseName: m.options.DatabaseName,
			Bookmarks:    queryParams.Bookmarks,
		})
		if err != nil {
			return nil, internalErr.ToDriverError(err)
		}
		defer usedSession.Close()

		// Run the query with the new session and the query config
		rawResult, err = usedSession.Run(queryParams.Query, paramsMap, queryParams.Configurers...)
		if err != nil {
			return nil, internalErr.ToDriverError(err)
		}
	}

	// Convert the raw results as a collection of typed results
	convertedResult := newQueryResult(rawResult)

	m.lastBookmark = usedSession.LastBookmark()

	return convertedResult, nil
}

func (m *manager) BeginTransaction(params TransactionParams) (string, internalErr.Neo4GoError) {
	usedSessionMode := neo4j.AccessModeRead
	if params.IsWrite {
		usedSessionMode = neo4j.AccessModeWrite
	}

	session, err := (*m.driver).NewSession(neo4j.SessionConfig{
		AccessMode:   usedSessionMode,
		DatabaseName: m.options.DatabaseName,
		Bookmarks:    params.Bookmarks,
	})
	if err != nil {
		return "", internalErr.ToDriverError(err)
	}

	tx, err := session.BeginTransaction(params.Configurers...)
	if err != nil {
		return "", internalErr.ToDriverError(err)
	}

	newTxUUID, err := uuid.NewV4()
	if err != nil {
		return "", &internalErr.TransactionError{
			Err: err.Error(),
		}
	}
	newTxID := newTxUUID.String()

	m.transactionSessions[newTxID] = transactionSession{
		session:     session,
		transaction: tx,
	}

	return newTxID, nil
}
func (m *manager) Commit(txID string) internalErr.Neo4GoError {
	txSession, exists := m.transactionSessions[txID]
	if !exists {
		return &internalErr.TransactionError{
			Err: "Trying to commit a non existing transaction",
		}
	}

	err := txSession.transaction.Commit()
	if err != nil {
		return internalErr.ToDriverError(err)
	}
	err = txSession.session.Close()
	if err != nil {
		return internalErr.ToDriverError(err)
	}

	delete(m.transactionSessions, txID)

	return nil
}
func (m *manager) Rollback(txID string) internalErr.Neo4GoError {
	txSession, exists := m.transactionSessions[txID]
	if !exists {
		return &internalErr.TransactionError{
			Err: "Trying to rollback a non existing transaction",
		}
	}

	err := txSession.transaction.Rollback()
	if err != nil {
		return internalErr.ToDriverError(err)
	}
	err = txSession.session.Close()
	if err != nil {
		return internalErr.ToDriverError(err)
	}

	delete(m.transactionSessions, txID)

	return nil
}

func (m *manager) LastBookmark() string {
	return m.lastBookmark
}
