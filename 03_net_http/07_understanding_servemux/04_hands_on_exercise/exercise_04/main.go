package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		io.WriteString(conn, "I see you connected.\n")
		conn.Close()
	}

	// go handle(listener)
}

// func handle(listener net.Listener) {
// 	defer listener.Close()
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		io.WriteString(conn, "I see you connected.\n")
// 		conn.Close()
// 	}
// }
