// Adaptado de Alan A. A. Donovan & Brian W. Kernighan.
// a TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "150.165.42.166:2000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {

	defer c.Close()
	for {
		_, err := io.WriteString(c, "SALVEEEE")

		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
