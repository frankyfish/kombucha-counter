package server

import (
	"log"
	"net/http"
)

// https://stackoverflow.com/a/31065973
var storage KombuchaStorage = NewRedisKombuchaStorage()

// todo: try with body handling.
func Start() {
	http.HandleFunc("/", getCurrentCount)
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
	w.Header().Add("Access-Control-Allow-Origin", "*") // CORS for simple request
	w.Write([]byte(*val))
}

func incCount(w http.ResponseWriter, r *http.Request) {
	log.Printf("Increasing. Request: %v\n", r)

	storage.IncCount(r.Context())
}
