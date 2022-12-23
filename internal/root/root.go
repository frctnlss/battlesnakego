package root

import (
	"encoding/json"
	"net/http"

	"battlesnake/m/v2/internal/pkg/battleSnake"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	default:
		http.NotFound(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(battleSnake.NewDefaultSnake())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
