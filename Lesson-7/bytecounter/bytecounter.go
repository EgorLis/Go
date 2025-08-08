package main

import (
	"fmt"
	"io"
)

type ByteCounter int

// реализуем контракт io.Writer
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

var _ io.Writer = (*ByteCounter)(nil)

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5
	c = 0          // сброс счетчика
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
}
