package main

import (
	"fmt"
)

// func main() {
// 	// var count = int(42)
// 	// ptr := &count
// 	// fmt.Println(*ptr)
// 	// *ptr = 100
// 	// fmt.Println(count)
// 	var guy = new(Person)
// 	guy.Name = "Shatyaki"
// 	// guy.sayHello()
// 	// greet(guy)
// 	// inter(guy)
// 	// normalForLoop()
// 	rangedForLoop()
// }

func normalForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("Printting ", i)
	}
}

func rangedForLoop() {
	nums := []int{2, 3, 4, 8}
	for _, val := range nums {
		fmt.Println("Index value is :", val)
	}
}

func inter(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("INteger value")
	default:
		fmt.Printf("Unknown value of type %T", v)
	}
}

func (p *Person) SayHello() {
	fmt.Println("Hello ,", p.Name)
}

type Friend interface {
	SayHello()
}

func greet(f Friend) {
	f.SayHello()
}

type Person struct {
	Name string
	Age  int
}
