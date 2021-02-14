package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	go print()
// 	// time.Sleep(3 * time.Second)
// 	// fmt.Println("main function")
// 	channel()
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("main function")
// }

func print() {
	fmt.Println("Pringing")
}

func channel() {
	c := make(chan int)
	go stringLength("Salutations", c)
	go stringLength("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	errors.New("SISSIH")
}

func stringLength(s string, c chan int) {
	c <- len(s)
}
func foo() error {
	return errors.New("Some Error Occurred")
}
