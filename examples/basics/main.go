package main

import (
	"log"

	"github.com/UlysseGuyon/neo4go/pkg/v1/neo4go"
	"github.com/mitchellh/mapstructure"
)

type User struct {
	Name string `neo4j:"name"`
}

func (u *User) ConvertToMap() map[string]neo4go.InputStruct {
	return map[string]neo4go.InputStruct{
		"name": neo4go.NewInputString(&u.Name),
	}
}

func main() {
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

	queryOpt := neo4go.QueryParams{
		Query: "WITH $newUser AS newU CREATE (u:User {name: newU.name}) RETURN u",
		Params: map[string]neo4go.InputStruct{
			"newUser": &User{Name: "Alice"},
		},
	}
	record, err := neo4go.Single(manager.Query(queryOpt))
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	userNode, exists := record.Nodes["u"]
	if !exists {
		log.Fatalln("u is null or not a node !")
	}
	user := User{}
	err = neo4go.NewNeo4GoDecoder(mapstructure.DecoderConfig{}).DecodeNode(userNode, &user)
	if err != nil {
		log.Fatalln(err.FmtError())
	}

	log.Printf("Saved user : %+v !", user)
}
