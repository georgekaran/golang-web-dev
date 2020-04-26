package main

import (
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

		handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	io.WriteString(conn, "I see you connected.\n")
	defer conn.Close()
}