package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// _, err := net.Dial("tcp", "scanme.nmap.org:80")
	// if err == nil {
	// 	fmt.Println("Connection successful")
	// }
	// scanMulti(1, 1024)
	ScanMulti(1, 1024)
}

func scanMulti(start int, end int) {
	var wg sync.WaitGroup
	for i := start; i < end; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			conn, err := Connect("tcp", fmt.Sprintf("scanme.nmap.org:%d", j))
			if err != nil {
				// port is closed or filtered.
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}

// Ecport
func Connect(connType string, addr string) (net.Conn, error) {
	return net.Dial(connType, addr)
}
