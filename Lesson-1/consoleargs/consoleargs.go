// Вывод аргументов командной строки
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Основная функция программы
func main() {
	var s, sep string

	start := time.Now()

	// := синтаксический сахар, нужен для присвоения локальных переменных

	// стандартный цикл
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	elapsed := time.Since(start)

	fmt.Printf("Итоговая строка 1: %s заняло времени: %d ms\n", s, elapsed)

	start = time.Now()
	s = ""
	sep = ""

	// цикл с помощью перечисляемых данных (в данном примере slice)
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	elapsed = time.Since(start)

	fmt.Printf("Итоговая строка 2: %s заняло времени: %d ms\n", s, elapsed)

	start = time.Now()

	s1 := strings.Join(os.Args[1:], " ")

	elapsed = time.Since(start)

	fmt.Printf("Итоговая строка 3: %s заняло времени: %d ms\n", s1, elapsed)

	fmt.Println(os.Args[1:])

	fmt.Println(os.Args)
}
