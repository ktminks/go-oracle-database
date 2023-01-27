# go-oracle-database

## Description

This is a simple Go program that connects to the CISE Oracle database and prints the results of a few queries.

### **Current Status: Semi-functional**

Known bug: Runs the first time successfully to completion, but subsequent runs fail with the following error:

> `query single row sql: Scan error on column index 1, name "CREATION_TIME": unsupported Scan, storing driver.Value type go_ora.TimeStamp into type *string`
>
> - go_ora is a package used to connect to the Oracle database
> - I believe the go_ora.TimeStamp type is getting confused with the form of TimeStamp that the database was expecting. Perhaps they share a name.
> - Could be that the code itself is wrong, as I am not very familiar with Go or OracleSQL yet.

## Installation

### Dependencies

> - [Oracle Instant Client](http://www.oracle.com/>technetwork/topics/linuxx86-64soft-092277.html)
>   - Must be properly set up. [(Instructions)]()
> - [golang](https://golang.org/dl/)
>   - GOPATH and/or GOROOT environment variables must be properly set. [(Instructions)](https://>golang.org/doc/install)
>   - Repo must be located within GOPATH/src/
> - other dependencies
>   - `go get .`

### Environment Variables

> To avoid hard-coding any sensitive information,store the credentials you use to connect to the CISE Oracle database in the following environment variables on your system:
>
> - UFL_USERNAME
> - ORACLE_PASSWORD

### Build & Run

> - `go run .`

## Credits

This project is based on the work of [lucasjellema](https://github.com/lucasjellema/go-oracle-database), who has kindly written a [tutorial](https://blogs.oracle.com/developers/post/connecting-a-go-application-to-oracle-database) explaining his code.  
I have made changes only to remove examples that are not relevant to our use case.
