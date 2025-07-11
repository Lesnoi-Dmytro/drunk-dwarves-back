package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJson(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

type errorReponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondJson(w, code, errorReponse{
		Code:  code,
		Error: msg,
	})
}
