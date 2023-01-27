package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/sijms/go-ora/v2"
)

var autonomousDB = map[string]string{
	"service":        "orcl",
	"username":       os.Getenv("UFL_USERNAME"),
	"server":         "oracle.cise.ufl.edu",
	"port":           "1521",
	"password":       os.Getenv("ORACLE_PASSWORD"),
}

func main() {
	var t = time.Now()

	doDBThings(autonomousDB)
	fmt.Println("Time Elapsed", time.Now().Sub(t).Milliseconds())
}

func handleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

const createTableStatement = "CREATE TABLE TEMP_TABLE ( NAME VARCHAR2(100), CREATION_TIME TIMESTAMP DEFAULT SYSTIMESTAMP, VALUE  NUMBER(5))"
const dropTableStatement = "DROP TABLE TEMP_TABLE PURGE"
const insertStatement = "INSERT INTO TEMP_TABLE ( NAME , VALUE) VALUES (:name, :value)"


func doDBThings(dbParams map[string]string) {
	connectionString := "oracle://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["service"]

	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

	someAdditionalActions(db)
}

func someAdditionalActions(db *sql.DB) {

	// var queryResultColumnOne string
	// row := db.QueryRow("SELECT systimestamp FROM dual")
	// err := row.Scan(&queryResultColumnOne)
	// if err != nil {
	// 	panic(fmt.Errorf("error scanning db: %w", err))
	// }
	// fmt.Println("The time in the database ", queryResultColumnOne)
	_, err := db.Exec(createTableStatement)
	handleError("create table", err)
	defer db.Exec(dropTableStatement)
	stmt, err := db.Prepare(insertStatement)
	handleError("prepare insert statement", err)
	sqlresult, err := stmt.Exec("John", 42)
	handleError("execute insert statement", err)
	rowCount, _ := sqlresult.RowsAffected()
	fmt.Println("Inserted number of rows = ", rowCount)

	var queryResultName string
	var queryResultTimestamp string
	var queryResultValue int32
	row := db.QueryRow("SELECT name, creation_time, value FROM temp_table")
	err = row.Scan(&queryResultName, &queryResultTimestamp, &queryResultValue)
	handleError("query single row", err)
	if err != nil {
		panic(fmt.Errorf("error scanning db: %w", err))
	}
	fmt.Println(fmt.Sprintf("The name: %s, time: %s, value:%d ", queryResultName, queryResultTimestamp, queryResultValue))

	_, err = stmt.Exec("Jane", 69)
	handleError("execute insert statement", err)
	_, err = stmt.Exec("Malcolm", 13)
	handleError("execute insert statement", err)

	// fetching multiple rows
	theRows, err := db.Query("select name, value from TEMP_TABLE")
	handleError("Query for multiple rows", err)
	defer theRows.Close()
	var (
		name  string
		value int32
	)
	for theRows.Next() {
		err := theRows.Scan(&name, &value)
		handleError("next row in multiple rows", err)
		fmt.Println(fmt.Sprintf("The name: %s and value:%d ", name, value))
	}
	err = theRows.Err()
	handleError("next row in multiple rows", err)

}
