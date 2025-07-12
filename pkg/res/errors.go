package res

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func Error(w http.ResponseWriter, code int, msg string) {
	Json(w, code, errorReponse{
		Code:  code,
		Error: msg,
	})
}

func BadRequestError(w http.ResponseWriter, msg string) {
	Error(w, 400, msg)
}

func MalformedJsonError(w http.ResponseWriter, err error) {
	Error(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %s", err))
}

func ValidationError(w http.ResponseWriter, err error) {
	errors := err.(validator.ValidationErrors)
	Error(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %s", errors))
}
