package types

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type ManagerOptions struct {
	URI          string
	DatabaseName string
	Realm        string
	Username     string
	Password     string
	CustomAuth   *neo4j.AuthToken
	Configurers  []func(*neo4j.Config)
}
