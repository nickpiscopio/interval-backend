package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
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
	//router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc(ENDPOINT_TIMER_STORE, storeTimer).Methods(METHOD_POST)
	router.HandleFunc(ENDPOINT_TIMER_GET, getTimer).Methods(METHOD_POST)

	log.Fatal(http.ListenAndServe(":" + PORT, router))
}