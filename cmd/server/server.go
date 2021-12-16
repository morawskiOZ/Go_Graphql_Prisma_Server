package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	graph "gitlab.com/morawskioz/camp_me_go/api/graphql"
	"gitlab.com/morawskioz/camp_me_go/api/graphql/generated"
	"gitlab.com/morawskioz/camp_me_go/cmd/server/middleware"
	"gitlab.com/morawskioz/camp_me_go/internal/prisma"
	"gitlab.com/morawskioz/camp_me_go/internal/prisma/db"
)

const defaultPort = "8080"

func main() {
	if err := prisma.LoadConfig(); err != nil {
		fmt.Printf("Fatal error: %+v\n", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	srv := handler.NewDefaultServer((generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Prisma: prismaClient}})))

	jwtMiddleware := middleware.CreateJWTMiddleware()

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", jwtMiddleware.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
