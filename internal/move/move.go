package move

import (
	"encoding/json"
	"log"
	"net/http"

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

type moveResponse struct {
	Move string `json:"move"`
}

func post(w http.ResponseWriter, r *http.Request) {
	payload := &battleSnake.Request{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("%+v\n", payload)
	result, err := json.Marshal(&moveResponse{Move: payload.Board.EdgeMovementClockwise(&payload.You.Head)})
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(result)
}
