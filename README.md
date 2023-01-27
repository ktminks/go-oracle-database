# go-oracle-database

## Description

This is a simple Go program that connects to the CISE Oracle database and prints the results of a few queries.  
It is meant to be a starting point for future projects that need to connect to the CISE Oracle database.

## Installation

### Environment Variables

#### GoLang

> For Go to compile and run, 2 environment variables [must be set properly](https://>golang.org/doc/install):
>
> - GOPATH: Root of your workspace.
>   - All Go code must be located within /GOPATH/src/
> - GOROOT: Location of Go SDK.
>   - Only necessary if you installed Go in a custom location.

#### Oracle CISE Database

> To avoid hard-coding any sensitive information into this repo, store the credentials you use to connect to the CISE Oracle database in the following environment variables on your system:
>
> - UFL_USERNAME
> - ORACLE_PASSWORD

### Prerequisites

> - [golang](https://golang.org/dl/)
> - [UFL VPN]https://net-services.ufl.edu/provided-services/vpn/clients/)

### Dependencies

> - [`database/sql`](https://golang.org/pkg/database/sql/): Go's standard database interface
> - [`go_ora`](github.com/sijms/go-ora/v2): connects to Oracle database
> - `time`: for time-related functions
> - `os`: reads environment variables
> - `fmt`: prints to console

### Build & Run

> - Download & install dependencies: `go get .`
> - Build & run: `go run .`

## More Resources

[Tutorial on Go database/sql](http://go-database-sql.org/overview.html) — general introduction to working with the database/sql package

[Conversation with sijms](https://github.com/sijms/go-ora/discussions/145) — creator of the go_ora library

[Introduction and overview the Oracle VS Code extension](https://www.oracle.com/database/technologies/appdev/dotnet/odtvscodequickstart.html)

## Credits

This project is based on the work of [lucasjellema](https://github.com/lucasjellema/go-oracle-database), who has kindly written a [tutorial](https://blogs.oracle.com/developers/post/connecting-a-go-application-to-oracle-database) explaining his code.  
I have made changes only to remove examples that are not relevant to our use case, and to simplify the code for others like myself who are new to Go.
