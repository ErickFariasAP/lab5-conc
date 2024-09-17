// Adaptado de Alan A. A. Donovan & Brian W. Kernighan.
// a TCP server that periodically writes the time.
package main

import (
	"encoding/binary"
	"fmt"
	"lab5-conc/config"
	"log"
	"net"
	"slices"
	"strings"
)

type ClientInfo struct {
	sum int64
	ip  string
}

var (
	mp         = make(map[int64][]string)
	registerCh = make(chan ClientInfo)
)

func register() {
	for {
		info := <-registerCh
		if info.sum > 0 {
			fmt.Printf("Received to Add: %d of %s\n", info.sum, info.ip)
			mp[info.sum] = append(mp[info.sum], info.ip)
		} else {
			fmt.Printf("Received to Remove: %d of %s\n", -info.sum, info.ip)
			mp[-info.sum] = slices.DeleteFunc(mp[-info.sum], func(s string) bool {
				return s == info.ip
			})
		}
	}
}

func registerServer() {
	listener, err := net.Listen("tcp", config.ServerIP+":"+config.RegisterPort)
	fmt.Println("Start register server")
	if err != nil {
		log.Fatal(err)
	}

	go register()
	for {
		conn, err := listener.Accept()
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

	for {
		var hash int64
		err := binary.Read(c, binary.BigEndian, &hash)
		if err != nil {
			fmt.Println("Error reading hash:", err)
			return
		}
		registerCh <- ClientInfo{hash, addr}
	}
}

func buscaServer() {
	for {
		listener, err := net.Listen("tcp", config.ServerIP+":"+config.SearchPort)
		fmt.Println("Start search server")
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := listener.Accept()
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

	var hash int64
	err := binary.Read(c, binary.BigEndian, &hash)
	if err != nil {
		fmt.Println("Error1 reading from connection:", err)
		return
	}

	fmt.Printf("Client %s searching for %d\n", addr, hash)

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
