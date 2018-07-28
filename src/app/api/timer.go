package api

import (
	"../structs"
	"../helper"
	"../helper/table"
	"encoding/json"
	"hash/fnv"
	"net/http"
	"time"
	"log"
)

/**
 * Gets the timer in the database.
 */
func GetTimer(rw http.ResponseWriter, req *http.Request) {
	var timer structs.Timer
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
 * Stores the timer in the database and returns the hash so users only need to use the hash as the url.
 */
func CreateTimer(rw http.ResponseWriter, req *http.Request) {
	var timer structs.Timer
	if req.Body == nil {
		http.Error(rw, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&timer)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	date := getTime()

	timer.Id = hash(timer.Timer)
	timer.DateCreated = date
	timer.DateUpdated = date
	timer.DateLastUsed = date

	helper.ExecuteStatement(table.GetInsert(), timer.Id, timer.Timer, timer.DateCreated, timer.DateUpdated, timer.DateLastUsed)

	json.NewEncoder(rw).Encode(timer)
}

/**
 * Updates the timer in the database and returns the hash so users only need to use the hash as the url.
 */
func UpdateTimer(rw http.ResponseWriter, req *http.Request) {
	var timer structs.Timer
	if req.Body == nil {
		http.Error(rw, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&timer)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	date := getTime()

	timer.DateUpdated = date
	timer.DateLastUsed = date

	helper.ExecuteStatement(table.GetPrepare(), timer.DateUpdated, timer.DateLastUsed)

	status := new(structs.Status)
	status.Code = 200
	status.Message = "Updated timer successfully."

	json.NewEncoder(rw).Encode(status)
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
