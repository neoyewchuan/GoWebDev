package main

import (
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	4

}

// to test this program, start main.go (go run main.go), then go to another
// terminal/shell and run the following command
// telnet localhost 8080
