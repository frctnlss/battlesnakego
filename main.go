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
	mux.HandleFunc("/start", startHandler)
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
	switch r.Method {
	case http.MethodGet:
		rootGet(w, r)
	default:
		http.NotFound(w, r)
	}
}

func rootGet(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(newDefaultSnake())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

type battleSnakeRequest struct {
	Game  game        `json:"game"`
	Turn  uint16      `json:"turn"`
	Board board       `json:"board"`
	You   battleSnake `json:"you"`
}

type game struct {
	Id      string  `json:"id"`
	Ruleset ruleset `json:"ruleset"`
	Map     string  `json:"map"`
	Timeout uint16  `json:"timeout"`
	Source  string  `json:"source"`
}

type ruleset struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type board struct {
	Height  uint8         `json:"height"`
	Width   uint8         `json:"width"`
	Food    []grid        `json:"food"`
	Hazards []grid        `json:"hazards"`
	Snakes  []battleSnake `json:"snakes"`
}

type grid struct {
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
}

type battleSnake struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Health         uint8          `json:"health"`
	Body           []grid         `json:"body"`
	Latency        string         `json:"latency"`
	Head           grid           `json:"head"`
	Length         uint16         `json:"length"`
	Shout          string         `json:"shout"`
	Squad          string         `json:"squad"`
	Customizations customizations `json:"customizations"`
}

type customizations struct {
	Color string `json:"color"`
	Head  string `json:"head"`
	Tail  string `json:"tail"`
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		startPost(w, r)
	default:
		http.NotFound(w, r)
	}
}

func startPost(w http.ResponseWriter, r *http.Request) {
	payload := &battleSnakeRequest{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
