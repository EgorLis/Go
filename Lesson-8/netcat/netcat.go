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

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // игнорируем ошибки
		log.Println("done")
		close(done)
		os.Stdin.Close()
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
	// ожидаем завершение фоновой подпрограммы
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
