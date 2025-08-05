package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Необходима передача двух аргументов!")
		return
	}

	if len(os.Args[1]) != len(os.Args[2]) {
		fmt.Println("Не совпадает количества символов в строках!")
		return
	}

	r1, r2 := []rune(os.Args[1]), []rune(os.Args[2])

	isAnagram := true
	length := len(r1)

	for i, c := range r1 {
		if c != r2[length-1-i] {
			isAnagram = false
			break
		}
	}

	s1, s2 := os.Args[1], os.Args[2]

	// операция удаление из среза

	b1 := []byte(s1)

	b1 = append(b1[:1], b1[2:]...)

	fmt.Println(string(b1))

	// --------------------------

	if isAnagram {
		fmt.Printf("%v обратное слову %v", s1, s2)
		return
	}

	fmt.Printf("%v не обратное слову %v", s1, s2)
}
