package server

import (
	"encoding/json"
	"log"
	"net/http"
)

const CORS_HEADER_NAME = "Access-Control-Allow-Origin"

// https://stackoverflow.com/a/31065973
var storage KombuchaStorage = NewRedisKombuchaStorage()

// todo: try with body handling.
func Start() {
	http.HandleFunc("/", getCurrentCount)
	http.HandleFunc("/stats", getCurrentStats)
	http.HandleFunc("/inc", incCount) // todo: limit to POST
	server := &http.Server{
		Addr: ":8080", //todo: make configurable
	}
	server.ListenAndServe()
}

func getCurrentCount(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %v\n", r)

	val, err := storage.GetCurrentCount(r.Context())

	if err != nil {
		panic(err)
	}
	w.Header().Add(CORS_HEADER_NAME, "*") // CORS for simple request
	w.Write([]byte(*val))
}

func getCurrentStats(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to get all stats: %v\n", r)

	val, err := storage.GetCurrentStats(r.Context())

	if err != nil {
		panic(err)
	}
	w.Header().Add(CORS_HEADER_NAME, "*") // CORS for simple request

	jsonResponse, err := json.Marshal(val)

	w.Write([]byte(*&jsonResponse))
}

func incCount(w http.ResponseWriter, r *http.Request) {
	log.Printf("Increasing. Request: %v\n", r)

	oh := r.Header["Origin"]
	if oh != nil {
		w.Header().Add(CORS_HEADER_NAME, oh[0])
	}

	storage.IncCount(r.Context())
}
