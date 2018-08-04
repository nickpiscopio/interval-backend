package helper

import (
	"./table"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const SQL_DATABASE = "sqlite3"
const DATABASE_NAME = "./Interval.db"

/**
 * Creates the database if it doesn't exist.
 */
func CreateDatabase() {
	_, error := ExecuteStatement(table.GetCreateTable())
	if error != nil {
		log.Println(error)
	}
}

/**
 * Executes a query.
 *
 * @param query		The query to execute.
 *
 * @return The result of the execution or the error that happened.
 */
func Execute(query string, args ...interface{}) (*sql.Rows, error) {
	database, error := connect()
	if error == nil {
		rows, error := database.Query(query, args...)
		if error == nil {
			database.Close()

			return rows, error
		} else {
			database.Close()

			return nil, error
		}
	} else {
		database.Close()

		return nil, error
	}
}

/**
 * Executes a statement.
 *
 * @param query		The query to execute.
 * @param args		The list of args to execute in the statement
 *
 * @return The result of the execution or the error that happened.
 */
func ExecuteStatement(query string, args ...interface{}) (sql.Result, error) {
	database, error := connect()
	if error == nil {
		statement, error := database.Prepare(query)
		if error == nil {
			result, error := statement.Exec(args...)

			database.Close()

			return result, error
		} else {
			database.Close()

			return nil, error
		}
	} else {
		database.Close()

		return nil, error
	}
}

/**
 * Connects to the database.
 *
 * @return The database connection or the error that happened.
 */
func connect() (*sql.DB, error) {
	return sql.Open(SQL_DATABASE, DATABASE_NAME)
}
