package main

import (
	"fmt"
	"os"
	"time"
)


func main() {
	t := time.Now()
	conn := Connect()
	fmt.Println("Time Elapsed", time.Since(t).Milliseconds())

	t = time.Now()
	RunTestQueries(conn)
	fmt.Println("Time Elapsed", time.Since(t).Milliseconds())

	defer Disconnect(conn)
}

func handleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
