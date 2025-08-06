package main

import "fmt"

// MustPositive не содержит return, но возвращает bool.
// Если x<0 — вернёт false, иначе — true.
func MustPositive(x int) (ok bool) {
	// Отложенный обработчик, который поймает panic и запишет
	// в ok значение, переданное в panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановилась")
			ok = r.(bool)
		}
	}()
	if x < 0 {
		panic(false) // «вернём» false
	}
	panic(true) // «вернём» true
}

func main() {
	fmt.Println(MustPositive(-10))

	// программа продолжит выполняться

	fmt.Println("End")

}
