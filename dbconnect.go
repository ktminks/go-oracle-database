package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/sijms/go-ora/v2"
)

var cred = map[string]string{
	"username":		os.Getenv("UFL_USERNAME"),
	"pw":		os.Getenv("ORACLE_PASSWORD"),
	"server":		"oracle.cise.ufl.edu",
	"port":			"1521",
	"sid":			"orcl",
}

func Connect() (**sql.DB) {
	connString := "oracle://" + cred["username"] + ":" + cred["pw"] +
		"@" + cred["server"] + ":" + cred["port"] + "/" + cred["sid"]
	
	db, err := sql.Open("oracle", connString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}
	return &db
}

func Disconnect(conn **sql.DB) () {
	db := *conn
	err := db.Close()
	if err != nil {
		fmt.Println("Can't close connection: ", err)
	}
}