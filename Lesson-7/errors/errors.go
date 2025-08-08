package main

import (
	"errors"
	"fmt"
)

type InvalidArgumentError struct {
	text string
}

func (e *InvalidArgumentError) Error() string { return e.text }
func (e *InvalidArgumentError) DoSmth() {
	fmt.Println("Особый обработчик для типа InvalidArgumentError")
}

type IndexOutOfRangeError struct {
	text string
}

func (e *IndexOutOfRangeError) Error() string { return e.text }
func (e *IndexOutOfRangeError) DoSmth() {
	fmt.Println("Особый обработчик для типа IndexOutOfRangeError")
}

func WhatIsError(e error) {
	switch e := e.(type) {
	case *InvalidArgumentError:
		fmt.Printf("[invalid_argument] Это ошибка типа: %T, со значением - %s\n", e, e)
		e.DoSmth()
	case *IndexOutOfRangeError:
		fmt.Printf("[out_of_range] Это ошибка типа: %T, со значением - %s\n", e, e)
		e.DoSmth()
	default:
		fmt.Printf("[default] Это ошибка типа: %T, со значением - %s\n", e, e)
	}
}

var _ error = (*IndexOutOfRangeError)(nil)
var _ error = (*InvalidArgumentError)(nil)

func main() {
	var err error
	err = &IndexOutOfRangeError{"значение вышло за пределы массива"}
	WhatIsError(err)
	err = &InvalidArgumentError{"нельзя присвоить символ числу"}
	WhatIsError(err)
	err = errors.New("EOF")
	WhatIsError(err)
}
