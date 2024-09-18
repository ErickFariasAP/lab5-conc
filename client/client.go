package main

import (
	"fmt"
	"lab5-conc/config"
	"net"
	"os"
)

var (
	old     = make(map[int64]bool)
	current = make(map[int64]bool)
)

func main() {
	var port string
	if os.Args[1] == "search" {
		port = config.SearchPort
	} else {
		port = config.UpdatePort
	}

	conn, err := net.Dial("tcp", config.ServerIP+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	if os.Args[1] == "search" {
		search(conn)
	} else {
		update(conn)
	}
}
