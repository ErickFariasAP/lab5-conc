package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
)

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

func sum(filePath string, c chan int64) {
	data, _ := readFile(filePath)

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	ans := _sum
	c <- int64(ans)
}

func sendNumToServer(conn net.Conn, x int64) {
	err := binary.Write(conn, binary.BigEndian, x)
	if err != nil {
		log.Fatal("Erro sending to Server: ", err)
	}
}

func calculateSums(c chan int64) {
	dir := os.Args[1]
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		go sum(dir+"/"+file.Name(), c)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		hsh := <-c
		current[hsh] = true
	}
}

func clearMap(mp map[int64]bool) {
	for k := range mp {
		delete(mp, k)
	}
}
