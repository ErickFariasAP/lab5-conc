// Adaptado de Alan A. A. Donovan & Brian W. Kernighan.
// a TCP server that periodically writes the time.
package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type ClientInfo struct {
	sum int64
	ip  string
}

var (
	mp         = make(map[int64][]string)
	registerCh = make(chan ClientInfo)
	mu         sync.Mutex
)

func register() {
	for {
		info := <-registerCh
		mu.Lock()
		mp[info.sum] = append(mp[info.sum], info.ip)
		mu.Unlock()
		fmt.Printf("Received: %d of %s\n", info.sum, info.ip)
	}
}

func registerServer() {
	listener, err := net.Listen("tcp", "localhost:2000")
	fmt.Println("Start register server")
	if err != nil {
		log.Fatal(err)
	}

	go register()
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

		registerCh <- ClientInfo{hash, addr}
	}
}

func buscaServer() {
	for {
		listener, err := net.Listen("tcp", "localhost:2001")
		fmt.Println("Start busca server")
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := listener.Accept()
			defer conn.Close()
			if err != nil {
				log.Print(err)
				continue
			}
			go handleSearchConn(conn)
		}

	}
}

func handleSearchConn(c net.Conn) {
	defer c.Close()
	addr := c.RemoteAddr().String()
	fmt.Printf("Client IP Address: %s\n", addr)

	var hash int64
	err := binary.Read(c, binary.BigEndian, &hash)
	if err != nil {
		fmt.Println("Error1 reading from connection:", err)
		return
	}
	ips := mp[hash]
	ips_string := strings.Join(ips, " ")

	ipsBytes := []byte(ips_string)
	ipsLen := int64(len(ipsBytes))

	err = binary.Write(c, binary.BigEndian, ipsLen)
	if err != nil {
		fmt.Println("Error writing IP length:", err)
		return
	}

	_, err = c.Write(ipsBytes)
	if err != nil {
		fmt.Println("Error writing IP string:", err)
		return
	}
}

func main() {
	go registerServer()
	go buscaServer()
	select {}
}
