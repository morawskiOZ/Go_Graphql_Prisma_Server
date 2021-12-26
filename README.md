# Camp Me Go API GQL
Graphql API for Camp Me App

### Technologies
#### Prisma Go Client
ORM client for postgres DB
#### gqlgen
Builds gql server

### Setup locally
1. Project needs `.env` in root folder and in `./internal/prisma`. Copy `.env.examples` in both respective folders as `.env` files and fill all variables.

2. When you first start or you have made changes to your model, migrate your database and re-generate your prisma code. From the root run:

    `make prepare`

It will prepare run migration and generate graphql server code

3. Run gql server

    `make dev`
