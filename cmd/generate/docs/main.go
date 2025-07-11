package main

import (
	"log"
	"os"

	"github.com/Lesnoi-Dmytro/drank-dwarves-api/pkg/exec"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	apiPath := os.Getenv("API_PATH")
	if apiPath == "" {
		log.Fatal("No API path defined in environment")
	}

	exec.ExecuteCommand("swag", "init", "-g", apiPath)
}
