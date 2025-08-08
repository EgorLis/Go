package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

const clientIDLETime = 10 * time.Second

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // например, обрыв соединения
			continue
		}
		wg.Add(1) // Добавили соединение
		go handleConn(conn, &wg)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	input := bufio.NewScanner(c)

	lines := make(chan string)

	go func() {
		defer close(lines)
		for input.Scan() {
			lines <- input.Text()
		}
	}()

	timer := time.NewTimer(clientIDLETime)
	defer timer.Stop()

	for {
		select {
		case line, ok := <-lines:
			if !ok {
				return
			}
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(clientIDLETime)
			go echo(c, line, 1*time.Second)
		case <-timer.C:
			fmt.Fprintf(c, "Время ожидания истекло...\n")
			c.Close()
			return
		}
	}
}
