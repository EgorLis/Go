package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	address := "localhost:8000"

	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
