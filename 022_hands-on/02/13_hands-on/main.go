package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main()  {
	listener, err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go serve(conn)
	}

}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	var i int
	var rMethod, rURI string
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
			break
		}
		i++
	}
	defer conn.Close()

	body := "Body..."
	body += "\n"
	body += rMethod
	body += "\n"
	body += rURI

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}