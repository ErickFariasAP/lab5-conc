package main

import (
	"encoding/binary"
	"fmt"
	"lab5-conc/config"
	"log"
	"net"
	"strings"
)

func searchServer() {
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
		fmt.Println("Error reading hash from connection:", err)
		return
	}

	fmt.Printf("Client %s searching for %d\n", addr, hash)

	ips := mp[hash]
	if len(ips) == 0 {
		ips = append(ips, "no-machines-with-this-hash")
	}

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
