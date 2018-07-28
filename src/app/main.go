package main

import (
	"./api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// This is the port that the server starts on.
// It needs to be a string.
const PORT = "8000"

const ENDPOINT_TIMER = "/timer"
const ENDPOINT_TIMER_STORE = ENDPOINT_TIMER + "/store"
const ENDPOINT_TIMER_GET = ENDPOINT_TIMER + "/get"

const METHOD_POST = "POST"

/**
 * This is the main function. It is where the application starts.
 */
func main() {
	log.Println("Started Interval!")

	router := mux.NewRouter()
	router.HandleFunc(ENDPOINT_TIMER_STORE, api.StoreTimer).Methods(METHOD_POST)
	router.HandleFunc(ENDPOINT_TIMER_GET, api.GetTimer).Methods(METHOD_POST)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
