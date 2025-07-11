package validation

import (
	"net/http"

	"github.com/Lesnoi-Dmytro/drank-dwarves-api/pkg/response"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateStruct[T interface{}](w http.ResponseWriter, data T) {
	err := validate.Struct(data)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		response.RespondWithError(w, 400, errors.Error())
	}
}
