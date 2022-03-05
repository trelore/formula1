## Formula GraphQL

Experimental play with GraphQL - needs major refactor

Goals:
- Transform the Formula 1 API into a graphQL API

### Running

To run this

```sh
go run server.go
```

and open a browser to localhost:8080.

Try a payload of 

```graphql
query {
  DriverStandings(filter: {year: "2018", top: 5}){
    drivers {
      points
      Driver {
        code
        givenName
        familyName
      }
    }
  }
}
```

### Hacking

Modify the `graph/schema.graphqls` file, and run `go run github.com/99designs/gqlgen generate`.
This should update any resolver function signatures, then modify the resolver function itself to support the new updates.

### Thoughts

GraphQL is incredible.
I can see many usecases of this, especially since it was so cheap to develop.
It compliments a classic REST API server really well.

# Agenda

- [x] Create local cluster
- [x] Deploy some services (formula1)
- [ ] Install Chaos Mesh
- [ ] Follow quick start guide