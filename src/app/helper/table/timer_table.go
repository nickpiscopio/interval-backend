package table

const TABLE_NAME = "TimerTable"

const COLUMN_ID = "id"
const COLUMN_TIMER = "timer"
const COLUMN_DATE_CREATED = "date_created"
const COLUMN_DATE_UPDATED = "date_updated"
const COLUMN_DATE_LAST_USED = "date_last_used"

/**
 * Gets the create the TimerTable query.
 *
 * @return The create table query.
 */
func GetCreateTable() string {
	return "CREATE TABLE IF NOT EXISTS " + TABLE_NAME + " (" + COLUMN_ID + " INTEGER PRIMARY KEY, " + COLUMN_TIMER + " TEXT, " + COLUMN_DATE_CREATED + " INTEGER, " + COLUMN_DATE_UPDATED + " INTEGER, " + COLUMN_DATE_LAST_USED + " INTEGER);"
}

/**
 * Gets the insert query to insert a timer into the database.
 *
 * @return The insert query.
 */
func GetInsert() string {
	return "INSERT INTO " + TABLE_NAME + " (" + COLUMN_ID + ", " + COLUMN_TIMER + ", " + COLUMN_DATE_CREATED + ", " + COLUMN_DATE_UPDATED + ", " + COLUMN_DATE_LAST_USED + ") VALUES (" + "?," + "?," + "?," + "?," + "?" + ");"
}

/**
 * Gets the prepare statement to retrieve a specified timer.
 *
 * @return The prepare statement.
 */
func GetTimerPrepare() string {
	return "SELECT " + COLUMN_TIMER + " FROM " + TABLE_NAME + " WHERE " + COLUMN_ID + " = ?;"
}

/**
 * Gets the prepare statement to update a specified timer.
 *
 * @param id			The Id of the timer to update.
 * @param withTimer		Boolean value on whether to include updating the timer or not.
 *
 * @return The prepare statement.
 */
func GetUpdatePrepare(withTimer bool) string {
	statement := "UPDATE " + TABLE_NAME + " SET " + COLUMN_DATE_UPDATED + " = ?, " + COLUMN_DATE_LAST_USED + " = ?"

	if withTimer {
		statement += ", " + COLUMN_TIMER + " = ?"
	}

	statement += "WHERE " + COLUMN_ID + " = ?;"
	return statement
}
