package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "joescatcam.website:80")
	if err != nil {
		log.Fatalln("Unable to get connection")
	}
	defer dst.Close()
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Inable to listen to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Inable to listen to port")
		}
		go handle(conn)
	}
}
