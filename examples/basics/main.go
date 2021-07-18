package main

import (
	"log"

	"github.com/UlysseGuyon/neo4go/pkg/v1/neo4go"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type User struct {
	Name string `neo4j:"name"`
}

func main() {
	// Instanciate the manager
	options := neo4go.ManagerOptions{
		URI:          "bolt://localhost",
		DatabaseName: "neo4j",
		Username:     "neo4j",
		Password:     "nPe4os45jWG0oroDdFloow",
		Configurers: []func(*neo4j.Config){
			func(c *neo4j.Config) {
				c.Encrypted = false
			},
		},
	}

	manager, err := neo4go.NewManager(options)
	if err != nil {
		log.Fatalln(err.FmtError())
	}
	defer manager.Close()

	// Create the query from encoded values
	encoder := neo4go.NewEncoder(nil)

	userAlice := User{Name: "Alice"}

	queryOpt := neo4go.QueryParams{
		Query: "WITH $newUser AS newU CREATE (u:User {name: newU.name}) RETURN u",
		Params: map[string]neo4go.InputStruct{
			"newUser": encoder.Encode(userAlice),
		},
	}

	// Run the query
	record, err := neo4go.Single(manager.Query(queryOpt))
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	// Decode the results from the query
	userNode, exists := record.Nodes["u"]
	if !exists {
		log.Fatalln("u is null or not a node !")
	}
	userRetreived := User{}
	err = neo4go.NewDecoder(nil).DecodeNode(userNode, &userRetreived)
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	log.Printf("Saved user : %+v !", userRetreived)
}
