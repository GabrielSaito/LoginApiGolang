package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCodde int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodde)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json: "error"`
	}{
		Error: err.Error(),
	})
}
