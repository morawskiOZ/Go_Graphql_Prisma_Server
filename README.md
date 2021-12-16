# Camp Me Go API GQL
Graphql API for Camp Me App

### Technologies
#### Prisma Go Client
ORM client for postgres DB
#### gqlgen
Builds gql server

### Setup

1. Whenever you make changes to your model, you should migrate your database and re-generate your prisma code:
```
# apply migrations
go run github.com/prisma/prisma-client-go migrate dev --name "init2"
# generate
go run github.com/prisma/prisma-client-go generate
```
2. Generate gql server code
`go run github.com/99designs/gqlgen generate`
