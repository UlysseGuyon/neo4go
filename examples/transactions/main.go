package main

import (
	"log"

	"github.com/UlysseGuyon/neo4go/pkg/v1/neo4go"
)

type User struct {
	Id   string `neo4j:"id"`
	Name string `neo4j:"name"`
}

func main() {
	// Instanciate the manager
	options := neo4go.ManagerOptions{
		URI:          "<YOUR_DATABASE_URI>",
		DatabaseName: "<YOUR_DATABASE_NAME>",
		Username:     "<YOUR_USERNAME>",
		Password:     "<YOUR_PASSWORD>",
	}

	manager, err := neo4go.NewManager(options)
	if err != nil {
		log.Fatalln(err.FmtError())
	}
	defer manager.Close()

	// Create the transaction from encoded values
	encoder := neo4go.NewEncoder(nil)
	decoder := neo4go.NewDecoder(nil)

	userAlice := User{Name: "Alice"}

	// Begin the transaction and save its ID
	txID, err := manager.BeginTransaction(neo4go.TransactionParams{IsWrite: true})
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	// Run the first query of the transaction by passing the txID as parameter
	firstQuery := neo4go.QueryParams{
		Query: "WITH $newUser AS newU CREATE (u:User {name: newU.name}) RETURN u",
		Params: map[string]neo4go.InputStruct{
			"newUser": encoder.Encode(userAlice),
		},
		Transaction: txID,
	}
	record, err := neo4go.Single(manager.Query(firstQuery))
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	userNode, exists := record.Nodes["u"]
	if !exists {
		log.Fatalln("User is null or not a node")
	}

	user := User{}
	err = decoder.DecodeNode(userNode, &user)
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	// Run the second query of the transaction like the first but tell it to commit at the end since it is the last query
	secondQuery := neo4go.QueryParams{
		Query: "MATCH (u:User {name: $userName}) SET u.id = 'abcd'RETURN u",
		Params: map[string]neo4go.InputStruct{
			"userName": neo4go.NewInputString(&user.Name),
		},
		Transaction:     txID,
		CommitOnSuccess: true,
	}
	record, err = neo4go.Single(manager.Query(secondQuery))
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	// Decode the result of the last query of the transaction
	userNode, exists = record.Nodes["u"]
	if !exists {
		log.Fatalln("u is null or not a node !")
	}
	userResult := User{}
	err = decoder.DecodeNode(userNode, &userResult)
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	log.Printf("Saved user : %+v !", userResult)
}
