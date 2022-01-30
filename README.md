# Go Gql Pirsma API
Graphql API with Prisma and Postgress starter. Example if how to connect those technologies to create fast and easy to use API.

### Technologies
- Graphql server, built with gqlgen
- postgress
- Prisma Go client, ORM for postgres DB

### Setup locally


1. Postgress DB is needed. For local env you can use docker compose config I prepared, run:

    `cd build && docker-compose up -d`
    
    Additionally you will have to create new db on you fresh postgres instance. You will use that db name in next step when filling env `DATABASE_URL="postgres://postgres@localhost:5432/{your_new_DB_name}"`. If you have your own postgress instance running, adjust the url accordingly.

1. Project needs `.env` in root folder and in `./internal/prisma`. Copy `.env.examples` in both respective folders as `.env` files and fill all variables.

1. When you first start or you have made changes to your model, migrate your database and re-generate your prisma code. From the root run:

    `make prepare`

    It will prepare run migration and generate graphql server code.

1. Run gql server

    `make dev`
