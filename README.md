# neo4go
Object oriented wrapper for the neo4j golang driver

[![CircleCI](https://circleci.com/gh/UlysseGuyon/neo4go.svg?style=svg)](https://circleci.com/gh/UlysseGuyon/neo4go)

Since the basic [neo4j-go-driver](https://github.com/neo4j/neo4j-go-driver#neo4j-go-driver) only allows you to pass a map of interfaces as parameters (but throws a runtime error if your map isn't correctly formatted) and only allows you to retreive interfaces as query results, all the type checking and convertion can become tedious as your app grows. The neo4go package simplifies all these processes.

## Installation

```bash
go get github.com/UlysseGuyon/neo4go
```

## Examples

- [Basics](examples/basics)
- [Transactions](examples/transactions)

## Getting Started

All the types and functions you should use are available in the `github.com/UlysseGuyon/neo4go/pkg/v1/neo4go` package. First, you need to instanciate a new neo4go manager.
```go
options := neo4go.ManagerOptions{
    URI: "<YOUR_DATABASE_URI>",
    DatabaseName: "<YOUR_DATABASE_NAME>",
    Username: "<YOUR_USERNAME>",
    Password: "<YOUR_PASSWORD>",
}

manager, err := neo4go.NewManager(options)
defer manager.Close()
```

Then, you can use the `Neo4GoEncoder` interface in order to convert your object as an input accepted by the neo4j-go-driver. When adding a specific tag (`neo4j` by default) to an exported struct field, you allow the encoder to find it and map it in the resulting object. All embeded value (in a field, a map or an array) that cannot be converted will be set to `nil` in the result object.

```go
// The struct we use as a neo4j node
type User struct {
    Name string `neo4j:"name"` // Here we add the `neo4j` tag to allow encoding
}

userAlice := User{Name: "Alice"}

queryOpt := neo4go.QueryParams{
    Query: "WITH $newUser AS newU CREATE (u:User {name: newU.name}) RETURN u",
    Params: map[string]neo4go.InputStruct{
        "newUser": neo4go.NewNeo4GoEncoder(nil).Encode(userAlice),
    },
}
record, err := neo4go.Single(manager.Query(queryOpt))
```

The result of a query is a list of maps, each containing typed objects. For example, if the result of your query is `RETURN 'abc' AS str`, then you should be able to access `str` through the following process.
```go
res, _ := manager.Query(neo4go.QueryParams{Query: "... RETURN 'abc' AS str"})

// str is already typed as a string
str, exists := res.Strings["str"]
if !exists {
    fmt.Println("'str' is either null or not a string at the end of the query !")
}
```

Nodes cannot be decoded directly in the result as the object type you want but you can still decode them easily thanks to the `Neo4GoDecoder` iterface. By adding the tag `neo4j` (or any other tag that you specify in the decoder options) for each node field that you want to map, you only need to pass an empty object and it will be filled automatically.

```go
type User struct {
	Name string `neo4j:"name"` // Here we added the `neo4j` tag to allow decoding
}

...

record, err := neo4go.Single(manager.Query("... RETURN u"))
if err != nil {
    log.Fatalln(err.FmtError())
}

userNode, exists := record.Nodes["u"]
if !exists {
    log.Fatalln("u is null or not a node !")
}
userRetreived := User{}
err = neo4go.NewNeo4GoDecoder(nil).DecodeNode(userNode, &userRetreived)
if err != nil {
    log.Fatalln(err.FmtError())
}

log.Printf("Saved user : %+v !", userRetreived)
```

## Licence

UlysseGuyon/neo4go is free and open-source software licensed under the [MIT License](LICENSE).