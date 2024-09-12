package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "150.165.42.166:2000")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send some data to the server
	for {
		tmp := make([]byte, 256)
		_, err := conn.Read(tmp)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(tmp))
	}

}
