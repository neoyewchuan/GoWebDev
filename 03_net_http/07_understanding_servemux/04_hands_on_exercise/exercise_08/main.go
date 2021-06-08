package main

import (
	"bufio"
	"fmt"
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
		// io.WriteString(conn, "I see you connected.\n")
		// conn.Close()
		go serve(conn)
	}

	// go handle(listener)
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		fmt.Println(ln)
	}

	fmt.Println("Code got here.")
	// Here we write to the response
	io.WriteString(conn, "I see you connected.")

}
