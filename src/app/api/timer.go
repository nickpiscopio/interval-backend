package api

import (
	"../constants"
	"../helper"
	"../helper/table"
	"../structs"
	"encoding/json"
	"hash/fnv"
	"net/http"
	"time"
)

/**
 * Gets the timer from it's ID.
 *
 * @return The timer from the ID that was specified or an error.
 */
func getTimer(id uint32) (string, error) {
	rows, error := helper.Execute(table.GetTimerPrepare(), id)
	if error == nil {
		var timer string
		// We defer the closing of the rows or else the database will return an error the next time we try to get the next row.
		// The error would be 'database is locked' if we didn't do this.
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&timer)

			return timer, nil
		}
	}

	return "", error
}

/**
 * Gets the timer in the database.
 */
func GetTimer(rw http.ResponseWriter, req *http.Request) {
	var timerStruct structs.Timer

	response := new(structs.Response)

	if req.Body == nil {
		response.Code = 400
		response.Body = constants.RESPONSE_NEED_BODY

		json.NewEncoder(rw).Encode(response)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&timerStruct)
	if err != nil {
		response.Code = 400
		response.Body = err.Error()

		json.NewEncoder(rw).Encode(response)
		return
	}

	timerVal, error := getTimer(timerStruct.Id)
	if error == nil {
		timer := new(structs.Timer)

		timer.Timer = timerVal

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

	// We want to see if a timer with that ID already exists first.
	timerVal, error := getTimer(timer.Id)
	if error == nil && len(timerVal) > 0 {
		// The timer already exists, so just send back the ID.
		// We do not create a new timer at this point in time.
		response.Code = 200
		response.Body = timer

		json.NewEncoder(rw).Encode(response)
	} else {
		// The timer doesn't exist, so we are going to store the one we received from request.
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
