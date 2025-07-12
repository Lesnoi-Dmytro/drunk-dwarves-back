package health

import (
	"net/http"

	"github.com/Lesnoi-Dmytro/drank-dwarves-api/pkg/res"
)

// GetHealth godoc
// @Summary      Health check
// @Description  Returns a 200 OK if the service is healthy
// @Tags         Health
// @Produce      application/json
// @Success      200 {object} ServerHealth
// @Router       /health [get]
func getHealthHandler(w http.ResponseWriter, r *http.Request) {
	health := getServerHealth()

	res.Json(w, 200, health)
}
