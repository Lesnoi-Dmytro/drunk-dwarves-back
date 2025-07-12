package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Lesnoi-Dmytro/drank-dwarves-api/.gen/drunk_dwarves/public/model"
	"github.com/Lesnoi-Dmytro/drank-dwarves-api/.gen/drunk_dwarves/public/table"
	_ "github.com/Lesnoi-Dmytro/drank-dwarves-api/docs"
	"github.com/Lesnoi-Dmytro/drank-dwarves-api/internal/config/api"
	"github.com/Lesnoi-Dmytro/drank-dwarves-api/internal/config/db"
	_ "github.com/Lesnoi-Dmytro/drank-dwarves-api/internal/config/db/migrations"
	"github.com/Lesnoi-Dmytro/drank-dwarves-api/internal/config/router"
	jet "github.com/go-jet/jet/v2/postgres"
	"github.com/joho/godotenv"
)

// @title Drunk Dwarves
// @version 1.0
// @description API written in Go for Drunk Dwarves Website
// @host localhost:8080
// @schemes http
// @BasePath /api
func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("No Port defined in environment")
	}

	log.Println("Starting server...")

	api.Config()

	srv := &http.Server{
		Handler: router.Create(),
		Addr:    ":" + portString,
	}

	user := model.Users{}
	query := jet.SELECT(table.Users.Email).
		FROM(table.Users.Table)
	err := query.Query(db.DB, &user)
	if err != nil {
		log.Println("Error retrieving user:", err)
	} else {
		log.Printf("%+v\n", user)
	}

	log.Printf("Server is running on port %s", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
