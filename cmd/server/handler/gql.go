package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	graph "gitlab.com/morawskioz/camp_me_go/api/graphql"
	"gitlab.com/morawskioz/camp_me_go/api/graphql/generated"
	"gitlab.com/morawskioz/camp_me_go/internal/prisma/db"
)

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GraphqlHandler() gin.HandlerFunc {
	prismaClient := db.NewClient()
	if err := prismaClient.Prisma.Connect(); err != nil {
		panic(err)
	}

	h := handler.NewDefaultServer((generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Prisma: prismaClient}})))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
