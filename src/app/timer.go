package main

import (
	"hash/fnv"
	"time"
	"net/http"
	"encoding/json"
	"log"
)

type Timer struct {
	Id uint32 `json:"id,omitempty"`
	Timer string `json:"timer"`
	StoredDate int64 `json:"date,omitempty"`
}

/**
 * Stores the timer in the database and returns the hash so users only need to use the hash as the url.
 */
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

	timer.Id = hash(timer.Timer)
	timer.StoredDate = getTime()

	json.NewEncoder(rw).Encode(timer)
}

/**
 * Stores the timer in the database and returns the hash so users only need to use the hash as the url.
 */
func getTimer(rw http.ResponseWriter, req *http.Request) {
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

	log.Println("Get timer!", timer.Id)

	// TODO: Get the timer from the database with the ID that was sent from the frontend.
	// TODO: This needs to be sent as a number
	//timer.Id

	json.NewEncoder(rw).Encode(timer)
}

/**
 * Hashes a string.
 *
 * @return The hashed string.
 */
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))

	return h.Sum32()
}

/**
 * Gets the time in milliseconds.
 *
 * @return The time in milliseconds.
 */
func getTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}