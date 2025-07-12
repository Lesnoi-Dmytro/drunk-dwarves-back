package feat

const HandlersFileTemplate = `package %[1]s

import (
	"net/http"

	"github.com/Lesnoi-Dmytro/drank-dwarves-api/pkg/res"
)

// Get%[2]s godoc
// @Summary      Get %[1]s
// @Description  Returns a list of %[1]s
// @Tags         %[2]s
// @Produce      application/json
// @Success      200 {object} %[3]sDto
// @Router       /%[1]s [get]
func get%[2]sHandler(w http.ResponseWriter, r *http.Request) {
	%[1]s := get%[2]s()

	res.Json(w, 200, %[1]s)
}
`
