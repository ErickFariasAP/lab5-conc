package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

type MyADT struct {
	sum  int
	path string
}

func sum(filePath string, c chan MyADT) {
	data, _ := readFile(filePath)

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	ans := MyADT{_sum, filePath}
	c <- ans
}

func main() {

	conn, err := net.Dial("tcp", "150.165.42.166:2000")
	if err != nil {
		fmt.Println(err)
		return
	}

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
