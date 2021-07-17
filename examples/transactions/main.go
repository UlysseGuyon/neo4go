package main

import (
	"errors"
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

	transactionParams := neo4go.TransactionParams{
		TransactionSteps: []neo4go.TransactionStepParams{
			{
				// First query of the transacction
				Query: "WITH $newUser AS newU CREATE (u:User {name: newU.name}) RETURN u",
				Params: map[string]neo4go.InputStruct{
					"newUser": encoder.Encode(userAlice),
				},
				TransitionFunc: func(qr neo4go.QueryResult) (map[string]neo4go.InputStruct, error) {
					// This function makes the transition between this query's result and the next query's params
					record, err := neo4go.Single(qr, nil)
					if err != nil {
						return nil, err
					}

					userNode, exists := record.Nodes["u"]
					if !exists {
						return nil, errors.New("User is null or not a node")
					}

					user := User{}
					err = decoder.DecodeNode(userNode, &user)
					if err != nil {
						return nil, err
					}

					return map[string]neo4go.InputStruct{
						"userName": neo4go.NewInputString(&user.Name),
					}, nil
				},
			},
			{
				Query: "MATCH (u:User {name: $userName}) SET u.id = 'abcd' RETURN u",
			},
		},
	}

	// Run the transaction
	record, err := neo4go.Single(manager.Transaction(transactionParams))
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	// Decode the result of the last query of the transaction
	userNode, exists := record.Nodes["u"]
	if !exists {
		log.Fatalln("u is null or not a node !")
	}
	user := User{}
	err = decoder.DecodeNode(userNode, &user)
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	log.Printf("Saved user : %+v !", user)
}
