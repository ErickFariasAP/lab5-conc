package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
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
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

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
