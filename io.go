package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}
type FooWriter struct{}

func (reader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in >")
	return os.Stdin.Read(b)
}

func (reader *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out <")
	return os.Stdin.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	input := make([]byte, 4096)
	// Use reader to read input.
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)
	// Use writer to write output.
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
	// if _, err := io.Copy(&writer, &reader); err != nil {
	// 	log.Fatalln("Unable to read data")
	// }
}
