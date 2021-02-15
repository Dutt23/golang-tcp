package main

import (
	"fmt"
	"net"
	"sort"
)

// Change number of workers and see
// Make the channel start sync, add print to see the results
func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go workers(ports, results)
	}
	// Seperate go routine. results channel needs to start before we 100 units of work is done
	go checkPorts(1024, ports)
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, index := range openports {
		fmt.Printf("%d open\n", index)
	}
}

func checkPorts(limit int, ports chan int) {
	for i := 0; i < limit; i++ {
		ports <- i
	}
}
func workers(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		// fmt.Println(err)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
