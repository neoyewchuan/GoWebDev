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
	var i int
	var rMethod, rURI string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
		fmt.Println(ln)
	}

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Hands-on Exercise #11</title></head>
	<body>
	<h1>HOLY COW THIS IS LOW LEVEL</h1><br>
	<h3>Method: ` + rMethod + `</h3><br>
	<h3>URI: ` + rURI + `</h3><br></body></html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "Request-Method: %s\r\n", rMethod)
	fmt.Fprintf(conn, "Request-URL: %s\r\n", rURI)
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}
