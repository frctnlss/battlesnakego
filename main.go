package main

import (
	"encoding/json"
	"log"
	"net/http"

	"battlesnake/m/v2/internal/root"
	"battlesnake/m/v2/internal/start"
	"battlesnake/m/v2/internal/end"
	"battlesnake/m/v2/internal/move"
)

func main() {
	addr := "0.0.0.0:8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", root.Handler)
	mux.HandleFunc("/start", start.Handler)
	mux.HandleFunc("/move", move.Handler)
	mux.HandleFunc("/end", end.Handler)
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
	isEmptyBody := r.Body == nil || r.Body == http.NoBody
	isApplicationJson := r.Header.Get("Content-Type") == "application/json"
	if !isEmptyBody && !isApplicationJson {
		http.Error(w, "Invalid Content-Type Header", http.StatusBadRequest)
		return
	}
	ej.handler.ServeHTTP(w, r)
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
