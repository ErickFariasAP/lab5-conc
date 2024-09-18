package main

type ClientInfo struct {
	sum int64
	ip  string
}

var (
	mp         = make(map[int64][]string)
	registerCh = make(chan ClientInfo)
)

func main() {
	go updateServer()
	go searchServer()
	select {}
}
