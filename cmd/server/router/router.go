package router

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/morawskioz/camp_me_go/cmd/server/handler"
	"gitlab.com/morawskioz/camp_me_go/cmd/server/middleware"
	"gitlab.com/morawskioz/camp_me_go/internal/prisma"

	"fmt"
)

// New sets up our routes and returns a *gin.Engine.
func New() *gin.Engine {
	if err := prisma.LoadConfig(); err != nil {
		fmt.Printf("Fatal error: %+v\n", err)
		os.Exit(1)
	}

	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:8080"},
			AllowCredentials: true,
			AllowHeaders:     []string{"Authorization"},
		},
	))

	router.GET(
		"/",
		handler.PlaygroundHandler(),
	)

	router.POST(
		"/query",
		middleware.EnsureValidToken(),
		handler.GraphqlHandler(),
	)

	return router
}
