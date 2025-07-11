package main

import (
	"log"
	"os"

	"github.com/Lesnoi-Dmytro/drank-dwarves-api/pkg/exec"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("No database URL defined in environment")
	}

	dbSchema := os.Getenv("DB_SCHEMA")
	if dbSchema == "" {
		log.Fatal("No database schema defined in environment")
	}

	exec.ExecuteCommand("jet", "-dsn="+dbUrl, "-schema="+dbSchema, "-path=./.gen")
}
