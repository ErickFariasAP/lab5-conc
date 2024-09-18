package main

import (
	"fmt"
	"net"
	"time"
)

func update(conn net.Conn) {
	c := make(chan int64)
	for {
		calculateSums(c)
		for k := range old {
			_, isPresent := current[k]

			if !isPresent {
				fmt.Printf("removing %d\n", k)
				sendNumToServer(conn, -k)
			}
		}

		for k := range current {
			_, isPresent := old[k]

			if !isPresent {
				fmt.Printf("adding %d\n", k)
				sendNumToServer(conn, k)
			}
		}

		clearMap(old)
		for k, v := range current {
			old[k] = v
		}
		clearMap(current)

		time.Sleep(time.Second)
	}
}
