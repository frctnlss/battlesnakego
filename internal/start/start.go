package start

import (
	"net/http"
	"encoding/json"
	"log"

	"battlesnake/m/v2/internal/pkg/battleSnake"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		post(w, r)
	default:
		http.NotFound(w, r)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	payload := &battleSnake.Request{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}