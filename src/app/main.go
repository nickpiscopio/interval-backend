package main

import (
	"fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"encoding/json"
)

/**
 * This is the main function. It is where the application starts.
 */
func main() {
	fmt.Println("INTERVAL BACKEND STARTING!")

    router := mux.NewRouter()
	//router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/timer", storeTimer).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8000", router))

    fmt.Println("INTERVAL BACKEND STARTED SUCCESSFULLY!")
}

type Timer struct {
	Id string `json:"id,omitempty"`
	Timer string `json:"timer"`
}

func storeTimer(rw http.ResponseWriter, req *http.Request) {
	var timer Timer
	if req.Body == nil {
		http.Error(rw, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&timer)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	fmt.Println(timer.Timer)

	timer.Id = "122"

	json.NewEncoder(rw).Encode(timer)
}