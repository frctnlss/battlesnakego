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
	handler := newContentTypeJson(newExpectJson(mux))
	log.Printf("Server running on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}

type expectJson struct {
	handler http.Handler
}

func newExpectJson(handler http.Handler) *expectJson {
	return &expectJson{handler}
}

func (ej *expectJson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil || r.Body == http.NoBody {
		ej.handler.ServeHTTP(w, r)
	} else if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type Header", http.StatusBadRequest)
	} else {
		ej.handler.ServeHTTP(w, r)
	}
}

type contentTypeJson struct {
	handler http.Handler
}

func newContentTypeJson(handler http.Handler) *contentTypeJson {
	return &contentTypeJson{handler}
}

func (ctj *contentTypeJson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	ctj.handler.ServeHTTP(w, r)
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
