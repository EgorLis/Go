package main

import "fmt"

func main() {
	p := 1
	b := inc(&p, 10)
	fmt.Printf("%T\n", &p)

	fmt.Printf("%v %v\n", &p, &b)

	// присваивания: = ++ --
	e := 1
	e++ // 2
	// присваивания кортежам
	c := 0
	e, c = c, e // 2 0 = 0 2
	fmt.Printf("e:%d c:%d\n", e, c)
}

func inc(p *int, delta int) int {
	fmt.Printf("%T\n", *p)
	fmt.Printf("%T\n", p)
	*p += delta
	return *p
}
