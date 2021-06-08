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

	switch {
	case rMethod == "GET" && rURI == "/":
		handleIndex(conn)
	case rMethod == "GET" && rURI == "/apply":
		handleApply(conn)
	case rMethod == "POST" && rURI == "/apply":
		handleApplyProcess(conn)
	}

}

func handleIndex(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Hands-on Exercise #12</title></head>
	<body><h1>INDEX</h1><br>
	<h4><a href="/">index</a></h4>
	<h4><a href="/apply">apply</a></h4></body></html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Request-Method: GET\r\n")
	fmt.Fprint(conn, "Request-URL: /\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Hands-on Exercise #12</title></head>
	<body><h1>APPLY</h1><br>
	<h4><a href="/">index</a></h4>
	<h4><a href="/apply">apply</a></h4>
	<form method="post" action="/apply">
	<input type="submit" value="apply">
	</form></body></html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Request-Method: GET\r\n")
	fmt.Fprint(conn, "Request-URL: /apply\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApplyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Hands-on Exercise #12</title></head>
	<body><h1>APPLY PROCESS</h1><br>
	<h4><a href="/">index</a></h4>
	<h4><a href="/apply">apply</a></h4></body></html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Request-Method: POST\r\n")
	fmt.Fprint(conn, "Request-URL: /apply\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
