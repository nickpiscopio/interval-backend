package table

const TABLE_NAME = "TimerTable"

const COLUMN_ID = "id"
const COLUMN_TIMER = "timer"
const COLUMN_DATE_CREATED = "date_created"
const COLUMN_DATE_UPDATED = "date_updated"
const COLUMN_DATE_LAST_USED = "date_last_used"

//statement1, _ := database.Prepare("insert into TimerTable(id, timer, date_created, date_updated, date_last_used) values (1, 'hello', 2, 3, 4);")


/**
 * Creates the TimerTable
 */
func CreateTable() string {
	//columns := make(map[string]string)
	//columns[COLUMN_ID] = "INTEGER PRIMARY KEY"
	//columns[COLUMN_TIMER] = "TEXT"
	//columns[COLUMN_DATE_CREATED] = "INTEGER"
	//columns[COLUMN_DATE_UPDATED] = "INTEGER"
	//columns[COLUMN_DATE_LAST_USED] = "INTEGER"

	//return createTable(TABLE_NAME, columns)

	return "CREATE TABLE IF NOT EXISTS " + TABLE_NAME + " (" + COLUMN_ID + " INTEGER PRIMARY KEY, " + COLUMN_TIMER + " TEXT, " + COLUMN_DATE_CREATED + " INTEGER, " + COLUMN_DATE_UPDATED + " INTEGER, " + COLUMN_DATE_LAST_USED + " INTEGER);"
}

func GetInsert() string {
	return "INSERT INTO " + TABLE_NAME + " (" + COLUMN_ID + ", " + COLUMN_TIMER + ", " + COLUMN_DATE_CREATED + ", " + COLUMN_DATE_UPDATED + ", " + COLUMN_DATE_LAST_USED + ") VALUES (" + "?," + "?," + "?," + "?," + "?" +");"
}

func GetPrepare() string {
	//columns := make(map[string]interface{})
	//columns[COLUMN_TIMER] = timer.Timer
	//columns[COLUMN_DATE_UPDATED] = timer.DateUpdated
	//columns[COLUMN_DATE_LAST_USED] = timer.DateLastUsed
	//log.Println("timer: ", timer)

	//query := "UPDATE " + TABLE_NAME + " SET " + COLUMN_TIMER + " = '" + timer.Timer + "', " + COLUMN_DATE_UPDATED + " = " + string(timer.DateUpdated) + ", " + COLUMN_DATE_LAST_USED + " = " + string(timer.DateLastUsed)

	return "UPDATE " + TABLE_NAME + " SET " + COLUMN_TIMER + " = ?, " + COLUMN_DATE_UPDATED + " = ?, " + COLUMN_DATE_LAST_USED + " = ?;"
}

/**
 * Builds and creates the query to create a table.
 *
 * @param tableName		The name of the table to create.
 * @param columns		The map of the columns with their data types that forms the table.
 */
//func createTable(tableName string, columns map[string]string) string {
//	query := "CREATE TABLE IF NOT EXISTS " + tableName + " ("
//
//	for column, dataType := range columns {
//		query += column + " " + dataType + ","
//	}
//
//	query += ");"
//
//	return query
//}

//func updateTable(tableName string, columns map[string]interface{}, whereColumn string, whereValue int32) string {
//	query := "UPDATE " + tableName + " SET "
//
//	for column, value := range columns {
//		query += column + " = "
//
//		switch value.(type) {
//			case int32:
//				query += string(value)
//			case string:
//				query += value
//			default:
//
//		}
//
//		query += ","
//	}
//
//	query += " WHERE " + whereColumn + " = " + whereValue + ";"
//
//	return query
//}