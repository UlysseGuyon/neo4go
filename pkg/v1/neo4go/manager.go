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

	// Query allows a single query to be made in database, possibly through an existing transaction
	Query(QueryParams) (QueryResult, Neo4GoError)

	// BeginTransaction starts a new transaction and stores it under the returned ID
	BeginTransaction(TransactionParams) (string, Neo4GoError)

	// Commit commits the transaction that has the given ID
	Commit(string) Neo4GoError

	// Rollback rolls back the transaction that has the given ID
	Rollback(string) Neo4GoError

	// LastBookmark returns the bookmark obtained by the session that ran the last query
	LastBookmark() string

	// Close closes the driver
	Close() Neo4GoError
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

	// The transaction ID to use for the query
	Transaction string

	// Whether to commit the transaction (if there is one for this query) after the successfull execution of the query
	CommitOnSuccess bool
}

// TransactionParams represents all the configuration of a single transaction at its creation
type TransactionParams struct {
	// Tells if the transaction type is read or write
	IsWrite bool

	// The bookmarks of previous sessions to apply to the query
	Bookmarks []string

	// The configurers to apply to the transaction
	Configurers []func(*neo4j.TransactionConfig)
}

// transactionSession represents a session and a transaction that are stored until the transaction is ended
type transactionSession struct {
	// The transaction that has been started
	transaction neo4j.Transaction

	// The session that started the transaction
	session neo4j.Session
}

// manager is the default implementation of the Manager interface
type manager struct {
	// The configuration of this manager
	options *ManagerOptions

	// The neo4j-go-driver wrapped in this manager
	driver *neo4j.Driver

	// The bookmark obtained by the session that ran the last query
	lastBookmark string

	// The map of transactions currently running, with their IDs as keys
	transactionSessions map[string]transactionSession
}

// NewManager creates a new instance of Manager, with a given config.
func NewManager(options ManagerOptions) (Manager, Neo4GoError) {
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
func (m *manager) init(options ManagerOptions) Neo4GoError {
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
		return toDriverError(err)
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
func (m *manager) Close() Neo4GoError {
	for _, val := range m.transactionSessions {
		err := val.transaction.Close()
		if err != nil {
			return toDriverError(err)
		}

		err = val.session.Close()
		if err != nil {
			return toDriverError(err)
		}
	}

	err := (*m.driver).Close()
	if err != nil {
		return toDriverError(err)
	}

	return nil
}

// Query allows a single query to be made in database, possibly through an existing transaction
func (m *manager) Query(queryParams QueryParams) (QueryResult, Neo4GoError) {
	// First, we convert all the input objects as interface maps
	paramsMap := make(map[string]interface{})
	for key, value := range queryParams.Params {
		paramsMap[key] = convertInputObject(value)
	}

	// Search an existing transaction with the given ID
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
		// If the transaction exists, then just run the query with it
		usedSession = txSession.session

		rawResult, err = txSession.transaction.Run(queryParams.Query, paramsMap)
		if err != nil {
			return nil, toDriverError(err)
		}

		if queryParams.CommitOnSuccess {
			err := m.Commit(queryParams.Transaction)
			if err != nil {
				return nil, err
			}
		}
	} else {
		// If the transaction does not exist, run the query as auto-commit transaction from a new session

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
			return nil, toDriverError(err)
		}
		defer usedSession.Close()

		// Run the query with the new session and the query config
		rawResult, err = usedSession.Run(queryParams.Query, paramsMap, queryParams.Configurers...)
		if err != nil {
			return nil, toDriverError(err)
		}
	}

	// Convert the raw results as a collection of typed results
	convertedResult := newQueryResult(rawResult)

	m.lastBookmark = usedSession.LastBookmark()

	return convertedResult, nil
}

// BeginTransaction starts a new transaction and stores it under the returned ID
func (m *manager) BeginTransaction(params TransactionParams) (string, Neo4GoError) {
	newTxUUID, err := uuid.NewV4()
	if err != nil {
		return "", &internalErr.TransactionError{
			Err: err.Error(),
		}
	}

	// First, init the session that will run the transaction
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
		return "", toDriverError(err)
	}

	// Then begin the transaction
	tx, err := session.BeginTransaction(params.Configurers...)
	if err != nil {
		closeErr := session.Close()
		if closeErr != nil {
			return "", toDriverError(closeErr)
		}
		return "", toDriverError(err)
	}

	// Finally, store the transaction and its session to the manager's map
	newTxID := newTxUUID.String()

	m.transactionSessions[newTxID] = transactionSession{
		session:     session,
		transaction: tx,
	}

	return newTxID, nil
}

// Commit commits the transaction that has the given ID
func (m *manager) Commit(txID string) Neo4GoError {
	// Get the transaction and its session from ID
	txSession, exists := m.transactionSessions[txID]
	if !exists {
		return &internalErr.TransactionError{
			Err: "Trying to commit a non existing transaction",
		}
	}

	// Commit the transaction and close its session
	err := txSession.transaction.Commit()
	if err != nil {
		return toDriverError(err)
	}
	err = txSession.session.Close()
	if err != nil {
		return toDriverError(err)
	}

	// Remove the transaction from the manager store
	delete(m.transactionSessions, txID)

	return nil
}

// Rollback rolls back the transaction that has the given ID
func (m *manager) Rollback(txID string) Neo4GoError {
	// Get the transaction and its session from ID
	txSession, exists := m.transactionSessions[txID]
	if !exists {
		return &internalErr.TransactionError{
			Err: "Trying to rollback a non existing transaction",
		}
	}

	// Commit the transaction and close its session
	err := txSession.transaction.Rollback()
	if err != nil {
		return toDriverError(err)
	}
	err = txSession.session.Close()
	if err != nil {
		return toDriverError(err)
	}

	// Remove the transaction from the manager store
	delete(m.transactionSessions, txID)

	return nil
}

// LastBookmark returns the bookmark obtained by the session that ran the last query
func (m *manager) LastBookmark() string {
	return m.lastBookmark
}
