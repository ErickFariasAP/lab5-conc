// Adaptado de Alan A. A. Donovan & Brian W. Kernighan.
// a TCP server that periodically writes the time.
package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

type ClientInfo struct {
	sum int64
	ip  string
}

var (
	mp    = make(map[int64][]string)
	canal = make(chan ClientInfo)
)

func cadastra() {
	for {
		info := <-canal
		mp[info.sum] = append(mp[info.sum], info.ip)
		fmt.Printf("Received: %d of %s\n", info.sum, info.ip)
	}
}

func registerServer() {
	listener, err := net.Listen("tcp", "localhost:2000")
	fmt.Println("Start register server")
	if err != nil {
		log.Fatal(err)
	}

	go cadastra()
	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleRegisterConn(conn)
	}
}

func handleRegisterConn(c net.Conn) {
	defer c.Close()
	addr := c.RemoteAddr().String()
	fmt.Printf("Client IP Address: %s\n", addr)

	for {
		var hash int64
		err := binary.Read(c, binary.BigEndian, &hash)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		canal <- ClientInfo{hash, addr}
	}
}

func main() {
	registerServer()
	//queryServer()
	select {}
}
