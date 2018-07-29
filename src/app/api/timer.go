package api

import (
	"../structs"
	"../helper"
	"../helper/table"
	"encoding/json"
	"hash/fnv"
	"net/http"
	"time"
	"../constants"
)

/**
 * Gets the timer in the database.
 */
func GetTimer(rw http.ResponseWriter, req *http.Request) {
	var timer1 structs.Timer

	response := new(structs.Response)

	if req.Body == nil {
		response.Code = 400
		response.Body = constants.RESPONSE_NEED_BODY

		json.NewEncoder(rw).Encode(response)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&timer1)
	if err != nil {
		response.Code = 400
		response.Body = err.Error()

		json.NewEncoder(rw).Encode(response)
		return
	}

	rows, error := helper.Execute(table.GetTimerPrepare(), timer1.Id)
	if err == nil {
		response.Code = 200

		var timer string
		for rows.Next() {
			rows.Scan(&timer)

			response.Body = timer
		}

		json.NewEncoder(rw).Encode(response)
	} else {
		response.Code = 400
		response.Body = error.Error()

		json.NewEncoder(rw).Encode(response)
	}
}

/**
 * Stores the timer in the database and returns the hash so users only need to use the hash as the url.
 */
func CreateTimer(rw http.ResponseWriter, req *http.Request) {
	var timer structs.Timer

	response := new(structs.Response)

	if req.Body == nil {
		response.Code = 400
		response.Body = constants.RESPONSE_NEED_BODY

		json.NewEncoder(rw).Encode(response)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&timer)
	if err != nil {
		response.Code = 400
		response.Body = err.Error()

		json.NewEncoder(rw).Encode(response)
		return
	}

	date := getTime()

	timer.Id = hash(timer.Timer)
	timer.DateCreated = date
	timer.DateUpdated = date
	timer.DateLastUsed = date

	_, error := helper.ExecuteStatement(table.GetInsert(), timer.Id, timer.Timer, timer.DateCreated, timer.DateUpdated, timer.DateLastUsed)
	if err == nil {
		response.Code = 200
		response.Body = timer

		json.NewEncoder(rw).Encode(response)
	} else {
		response.Code = 400
		response.Body = error.Error()

		json.NewEncoder(rw).Encode(response)
	}
}

/**
 * Updates the timer in the database and returns the hash so users only need to use the hash as the url.
 */
func UpdateTimer(rw http.ResponseWriter, req *http.Request) {
	var timer structs.Timer

	response := new(structs.Response)

	if req.Body == nil {
		response.Code = 400
		response.Body = constants.RESPONSE_NEED_BODY

		json.NewEncoder(rw).Encode(response)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&timer)
	if err != nil {
		response.Code = 400
		response.Body = err.Error()

		json.NewEncoder(rw).Encode(response)
		return
	}

	date := getTime()

	timer.DateUpdated = date
	timer.DateLastUsed = date
	updateWithTimer := len(timer.Timer) > 0

	var args []interface{}
	args = append(args, timer.DateUpdated)
	args = append(args, timer.DateLastUsed)

	if updateWithTimer {
		args = append(args, timer.Timer)
	}

	args = append(args, timer.Id)

	_, error := helper.ExecuteStatement(table.GetUpdatePrepare(updateWithTimer), args...)
	if err == nil {
		response.Code = 200
		response.Body = constants.RESPONSE_UPDATED_SUCCESSFULLY

		json.NewEncoder(rw).Encode(response)
	} else {
		response.Code = 400
		response.Body = error.Error()

		json.NewEncoder(rw).Encode(response)
	}
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
