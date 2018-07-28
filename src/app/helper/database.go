package helper

import (
	"./table"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"log"
)

const DATABASE_NAME = "./Interval.db"

/**
 * Creates the database if it doesn't exist.
 */
func CreateDatabase() {
	Execute(table.CreateTable())

	//rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	//var id int
	//var firstname string
	//var lastname string
	//for rows.Next() {
	//	rows.Scan(&id, &firstname, &lastname)
	//	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	//}
}

/**
 * Executes a query.
 *
 * @param query		The query to execute.
 */
func Execute(query string) {
	database, _ := connect()
	statement, _ := database.Prepare(query)
	statement.Exec()

	database.Close()
}

/**
 * Executes a statement.
 *
 * @param query		The query to execute.
 * @param args		The list of args to execute in the statement
 */
func ExecuteStatement(query string, args interface{}) {
	log.Println("query: ", query)
	log.Println("args: ", args)
	database, _ := connect()
	statement, _ := database.Prepare("INSERT INTO TimerTable (id, timer, date_created, date_updated, date_last_used) VALUES (?,?,?,?,?);")
	//statement, _ := database.Prepare(query)
	log.Println("statement: ", statement)
	log.Println("database: ", database)
	//statement.Exec(933223032, "1234a568",1532810406164,1532810406164,1532810406164)
	result, error := statement.Exec(args)

	log.Println("result: ", result)
	log.Println("error: ", error)

	database.Close()
}

/**
 * Connects to the database.
 */
func connect() (*sql.DB, error) {
	return sql.Open("sqlite3", DATABASE_NAME)
}
