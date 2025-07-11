package response

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func MalformedJson(w http.ResponseWriter, err error) {
	RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
}

func ValidationFailed(w http.ResponseWriter, err error) {
	errors := err.(validator.ValidationErrors)
	RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", errors))
}
