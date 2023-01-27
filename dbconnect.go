package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/sijms/go-ora/v2"
)

var dbParams = map[string]string{
	"sid":        "orcl",
	"username":       os.Getenv("UFL_USERNAME"),
	"server":         "oracle.cise.ufl.edu",
	"port":           "1521",
	"password":       os.Getenv("ORACLE_PASSWORD"),
}

func Connect() {
	connectionString := "oracle://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["sid"] + "?parseTime=true"

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

	
	var t = time.Now()
	RunTestQueries(db)
	fmt.Println("Time Elapsed", time.Since(t).Milliseconds())
}