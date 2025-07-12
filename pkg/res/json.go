package res

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, code int, payload any) {
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

func SuccessJson(w http.ResponseWriter, payload any) {
	Json(w, 200, payload)
}

type errorReponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
