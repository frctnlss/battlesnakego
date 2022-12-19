package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	addr := "0.0.0.0:8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	log.Printf("Server running on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

type snake struct {
	ApiVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

func newDefaultSnake() *snake {
	return &snake{
		ApiVersion: "1",
		Author:     "frctnlss",
		Color:      "#7CFF37",
		Head:       "missile",
		Tail:       "missile",
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(newDefaultSnake())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
