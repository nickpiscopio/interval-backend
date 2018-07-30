package main

import (
	"./api"
	"./helper"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// This is the port that the server starts on.
// It needs to be a string.
const PORT = "8000"

const ENDPOINT_TIMER = "/timer"
const ENDPOINT_TIMER_CREATE = ENDPOINT_TIMER + "/create"
const ENDPOINT_TIMER_UPDATE = ENDPOINT_TIMER + "/update"
const ENDPOINT_TIMER_GET = ENDPOINT_TIMER + "/get"

const METHOD_POST = "POST"

func main() {
	log.Println("Started Interval!")

	helper.CreateDatabase()

	router := mux.NewRouter()
	router.HandleFunc(ENDPOINT_TIMER_CREATE, api.CreateTimer).Methods(METHOD_POST)
	router.HandleFunc(ENDPOINT_TIMER_UPDATE, api.UpdateTimer).Methods(METHOD_POST)
	router.HandleFunc(ENDPOINT_TIMER_GET, api.GetTimer).Methods(METHOD_POST)
	http.Handle("/", &MyServer{router})

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}
