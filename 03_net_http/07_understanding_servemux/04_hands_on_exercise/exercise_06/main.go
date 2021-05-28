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

		//defer conn.Close()

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			if ln == "" {
				break
			}
			fmt.Println(ln)
		}

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")
		conn.Close()
	}

	// go handle(listener)
}
