migrate:
	cd internal/prisma && go run github.com/prisma/prisma-client-go migrate dev
generate:
	cd internal/prisma && go run github.com/prisma/prisma-client-go generate
prepare: migrate generate
dev:
	cd cmd/server && go run server.go