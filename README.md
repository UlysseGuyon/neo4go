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

Then, for each type that you want to pass as a neo4j query parameter, you first have to implement the following method.
```go
// The return value is a mapping of each field name/value if you type is a struct.
ConvertToMap() map[string]InputStruct
```
If you are using a type that cannot be mapped (like an array), your `ConvertToMap` method should return nil and you should also implement the following method.
```go
// You can either return a struct or a primitive value.
ConvertToInputObject() InputStruct
```
For each primitive type that can be found in neo4j, you can instanciate the equivalent as a neo4go `InputStruct` with provided functions such as `NewInputInteger`, `NewInputString`, `NewInputArray`, etc.

Then, you can call any query you want with these parameters.
```go
// The struct we use as a neo4j node
type User struct {
    Name string
}

// Implementation of ConvertToMap in order to convert `User` as an `InputStruct`
func (u *User) ConvertToMap() map[string]neo4go.InputStruct {
    return map[string]neo4go.InputStruct{
        "name": neo4go.NewInputString(&u.Name),
    }
}

queryOpt := neo4go.QueryParams{
	Query: "WITH $newUser AS newU CREATE (u:User {name: newU.name})",
	Params: map[string]neo4go.InputStruct{
        "newUser": &User{Name: "Alice"},
    },
}
res, err := manager.Query(queryOpt)
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
Nodes cannot be decoded directly in the result as the object type you want but you can still decode them easily thanks to the `DecodeNode` function. By adding the tag `neo4j` (or any other tag that you specify in the decoder options) for each node field that you want to map, you only to pass an empty object and it will be filled automatically.
```go
type User struct {
	Name string `neo4j:"name"` // Here we added the `neo4j` tag
}

...

record, _ := neo4go.Single(manager.Query(neo4go.QueryParams{Query: "... RETURN user"})) // user is a node

userNode, exists := record.Nodes["user"]
if !exists {
    fmt.Println("user is null or not a node !")
}
user := User{}
err = neo4go.NewNeo4GoDecoder(mapstructure.DecoderConfig{}).DecodeNode(userNode, &user) // you should give a pointer so that it can be modified
if err != nil {
    log.Fatalln(err.FmtError())
}

log.Printf("User : %+v", user)
```

## Licence

UlysseGuyon/neo4go is free and open-source software licensed under the [MIT License](LICENSE).