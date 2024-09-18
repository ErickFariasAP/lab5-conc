package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func search(conn net.Conn) {
	hsh, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	sendNumToServer(conn, int64(hsh))
	var ipLen int64
	err = binary.Read(conn, binary.BigEndian, &ipLen)
	if err != nil {
		fmt.Println("Error reading string length from connection:", err)
		return
	}

	ipBytes := make([]byte, ipLen)

	_, err = conn.Read(ipBytes)
	if err != nil {
		fmt.Println("Error reading string from connection:", err)
		return
	}

	ip := string(ipBytes)

	fmt.Println(strings.ReplaceAll(ip, " ", "\n"))
}
