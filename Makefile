migrate:
	cd internal/prisma && go run github.com/prisma/prisma-client-go migrate dev
generate-prisma:
	cd internal/prisma && go run github.com/prisma/prisma-client-go generate
generate-gql:
	cd api && go run github.com/99designs/gqlgen generate
prepare: generate-gql generate-prisma migrate
dev:
	cd cmd/server && go run server.go