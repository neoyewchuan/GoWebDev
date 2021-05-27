package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)

	}

}

func handle(conn net.Conn) {

	defer conn.Close()

	// instructions
	io.WriteString(conn, "\nIN-MEMORY DATABASE\n\n"+
		"USE:\n"+
		"SET key value \n"+
		"GET key \n"+
		"DEL key \n"+
		"EXAMPLE: \n"+
		"SET fav chocolate \n"+
		"GET fav \n\n\n")

	// read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		// logic
		switch fs[0] {
		case "GET":
			if len(fs) != 2 {
				fmt.Fprintln(conn, "EXPECTED KEY")
				continue
			}
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			if len(fs) != 2 {
				fmt.Fprintln(conn, "EXPECTED KEY")
				continue
			}
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND")
		}

	}

}

// You can test this code using telnet localhost 8080 from another terminal/shell
