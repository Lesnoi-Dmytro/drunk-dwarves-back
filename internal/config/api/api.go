package api

import (
	"github.com/Lesnoi-Dmytro/drank-dwarves-api/internal/config/db"
)

func Config() {
	db.ConnectDB()
}
