package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func echo(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		s, err := reader.ReadString(('\n'))

		if err == io.EOF {
			log.Println("Client has disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error has occured")
			break
		}

		if len(strings.TrimSpace(s)) <= 0 {
			continue
		}

		log.Printf("Received %d bytes: %s\n", len(s), s)
		// Send data via conn.Write.
		log.Println("Writing data")
		writer := bufio.NewWriter(conn)
		if _, err := writer.Write([]byte(s)); err != nil {
			log.Fatalln("Unable to write data")
		}
		// Required to flush the dat into the socket for reply
		writer.Flush()

		// One line method
		// if _, err := io.Copy(conn, conn); err != nil {
		// 	log.Fatalln("Unable to read or write data")
		// }
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Println("Unable to bind port")
	}

	log.Println("Listening on port 20080")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept conection")
		}
		go echo(conn)
	}
}
