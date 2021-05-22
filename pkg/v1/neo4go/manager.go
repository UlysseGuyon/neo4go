package neo4go

import (
	"github.com/UlysseGuyon/neo4go/internal/errors"
	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	internalTypes "github.com/UlysseGuyon/neo4go/internal/types"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type ManagerOptions struct {
	URI          string
	DatabaseName string
	Username     string
	Password     string
	CustomAuth   *neo4j.AuthToken
	Configurers  []func(*neo4j.Config)
}

func validateManagerOptions(opt ManagerOptions) internalErr.Neo4GoError {
	if opt.URI == "" {
		return internalErr.Neo4GoInitError{
			Bare:   true,
			Reason: "Database URI given in options is empty",
		}
	}

	if opt.DatabaseName == "" {
		return internalErr.Neo4GoInitError{
			Bare:   true,
			Reason: "Database name given in options is empty",
		}
	}

	return nil
}

type BasicQueryParams struct {
	Query   string
	Params  map[string]interface{}
	IsWrite bool
}

type QueryParams struct {
	BasicQueryParams
	ExpectedResult []internalTypes.ExpectedResult
}

type Manager interface {
	Init(ManagerOptions) internalErr.Neo4GoError
	IsConnected() bool
	Close()
	BasicQuery(BasicQueryParams) (neo4j.Result, internalErr.Neo4GoError)
	Query(QueryParams) internalErr.Neo4GoError
}

func NewManager(options ManagerOptions) (Manager, internalErr.Neo4GoError) {
	newManager := manager{}

	err := newManager.Init(options)
	if err != nil {
		return nil, err
	}

	if !newManager.IsConnected() {
		return nil, errors.Neo4GoConnError{
			URI:    options.URI,
			DBName: options.DatabaseName,
		}
	}

	return &newManager, nil
}

type manager struct {
	options      *ManagerOptions
	Driver       *neo4j.Driver
	SessionRead  *neo4j.Session
	SessionWrite *neo4j.Session
}

func (m manager) Init(options ManagerOptions) internalErr.Neo4GoError {
	optErr := validateManagerOptions(options)
	if optErr != nil {
		return optErr
	}

	usedAuth := neo4j.NoAuth()
	if options.CustomAuth != nil {
		usedAuth = *options.CustomAuth
	} else if options.Username != "" && options.Password != "" {
		usedAuth = neo4j.BasicAuth(options.Username, options.Password, "")
	}

	newDriver, err := neo4j.NewDriver(
		options.URI,
		usedAuth,
		options.Configurers...,
	)

	if err != nil {
		return internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: err.Error(),
		}
	}

	sessionWrite, err := newDriver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: options.DatabaseName,
	})

	if err != nil {
		return internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: err.Error(),
		}
	}

	sessionRead, err := newDriver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: options.DatabaseName,
	})

	if err != nil {
		return internalErr.Neo4GoInitError{
			Bare:   false,
			Reason: err.Error(),
		}
	}

	m.options = &options
	m.Driver = &newDriver
	m.SessionWrite = &sessionWrite
	m.SessionRead = &sessionRead

	return nil
}

func (m manager) IsConnected() bool {
	if m.Driver == nil {
		return false
	}

	err := (*m.Driver).VerifyConnectivity()

	return err == nil
}

func (m manager) Close() {
	if m.SessionRead != nil {
		_ = (*m.SessionRead).Close()
	}

	if m.SessionWrite != nil {
		_ = (*m.SessionWrite).Close()
	}

	if m.Driver != nil {
		_ = (*m.Driver).Close()
	}
}

func (m manager) BasicQuery(params BasicQueryParams) (neo4j.Result, internalErr.Neo4GoError) {
	if params.IsWrite && m.SessionWrite == nil {
		return nil, internalErr.Neo4GoQueryBuildError{
			Bare:   true,
			Reason: "Write Session is not initialised in the manager",
		}
	} else if params.IsWrite && m.SessionWrite != nil {
		res, err := (*m.SessionWrite).Run(params.Query, params.Params)
		if err != nil {
			return nil, internalErr.Neo4JQueryError{
				Bare:   false,
				Reason: err.Error(),
			}
		}

		return res, nil
	}

	if !params.IsWrite && m.SessionRead == nil {
		return nil, internalErr.Neo4GoQueryBuildError{
			Bare:   true,
			Reason: "Read Session is not initialised in the manager",
		}
	} else if !params.IsWrite && m.SessionRead != nil {
		res, err := (*m.SessionRead).Run(params.Query, params.Params)
		if err != nil {
			return nil, internalErr.Neo4JQueryError{
				Bare:   false,
				Reason: err.Error(),
			}
		}

		return res, nil
	}

	return nil, internalErr.Neo4JUnknownError{}
}

func (m manager) Query(params QueryParams) internalErr.Neo4GoError {
	// TODO implement
	return nil
}
