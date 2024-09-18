package main

import (
	"encoding/binary"
	"fmt"
	"lab5-conc/config"
	"log"
	"net"
	"slices"
)

func update() {
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

func updateServer() {
	listener, err := net.Listen("tcp", config.ServerIP+":"+config.UpdatePort)
	fmt.Println("Start register server")
	if err != nil {
		log.Fatal(err)
	}

	go update()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleUpdateConn(conn)
	}
}

func HandleUpdateConn(c net.Conn) {
	defer c.Close()
	addr := c.RemoteAddr().String()

	for {
		var hash int64
		err := binary.Read(c, binary.BigEndian, &hash)
		if err != nil {
			fmt.Println("Error reading hash", err)
			return
		}
		registerCh <- ClientInfo{hash, addr}
	}
}
