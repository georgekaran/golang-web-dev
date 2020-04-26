package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
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

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You said: %v.\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.\n")
}