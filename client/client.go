package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

func sum(filePath string, c chan int) {
	data, _ := readFile(filePath)

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	ans := _sum
	c <- ans
}

func main() {
	var port string
	if len(os.Args) > 1 && os.Args[1] == "search" {
		port = "2001"
	} else {
		port = "2000"
	}

	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	if os.Args[1] == "search" {
		hsh, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		binary.Write(conn, binary.BigEndian, int64(hsh))
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
	} else {
		c := make(chan int)
		for _, path := range os.Args[1:] {
			go sum(path, c)
		}

		for range os.Args[1:] {
			hsh := <-c
			fmt.Println(hsh)
			binary.Write(conn, binary.BigEndian, int64(hsh))
		}
	}
}
