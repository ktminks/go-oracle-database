package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var t = time.Now()
	Connect()
	fmt.Println("Time Elapsed", time.Since(t).Milliseconds())
}

func handleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
