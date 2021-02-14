package main

import (
	"fmt"
	"sync"
)

func ScanMulti(start int, end int) {
	var wg sync.WaitGroup
	for i := start; i < end; i++ {
		// 	conn, err := connect("tcp", fmt.Sprintf("scanme.nmap.org:%d", i))
		// 	if err != nil {
		// 		// port is closed or filtered.
		// 		continue
		// 	}
		// 	conn.Close()
		// 	fmt.Printf("%d open\n", i)
		// }
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
