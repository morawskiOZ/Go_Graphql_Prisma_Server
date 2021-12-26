package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/morawskioz/camp_me_go/cmd/server/router"
	"gitlab.com/morawskioz/camp_me_go/internal/prisma"
)

const defaultPort = "8080"

func main() {
	if err := prisma.LoadConfig(); err != nil {
		fmt.Printf("Fatal error: %+v\n", err)
		os.Exit(1)
	}

	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Printf("Fatal error: %+v\n", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	rtr := router.New()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, rtr))

}
