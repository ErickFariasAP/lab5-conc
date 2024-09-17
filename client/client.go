package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	old  = make(map[int64]bool)
	novo = make(map[int64]bool)
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
		c := make(chan int64)
		for {
			//fmt.Print("CHECK\n")
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
				//	fmt.Printf("%d\n", hsh)
				novo[hsh] = true
			}

			for k := range old {
				_, ok := novo[k]

				if !ok {
					fmt.Printf("removing %d\n", k)

					binary.Write(conn, binary.BigEndian, -k)
				}
			}

			for k := range novo {
				_, ok := old[k]

				if !ok {
					binary.Write(conn, binary.BigEndian, k)
					fmt.Printf("sending %d\n", k)
				}
			}

			for k := range old {
				delete(old, k)
			}

			for k, v := range novo {
				old[k] = v
			}

			for k := range novo {
				delete(novo, k)
			}

			time.Sleep(time.Second)
		}
	}
}
