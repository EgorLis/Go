package main

import (
	"examples/Lesson-9/bank"
	"fmt"
	"sync"
)

func main() {
	var n sync.WaitGroup

	n.Add(1)
	go func() {
		defer n.Done()
		bank.Deposit(200)                // A1
		fmt.Println("=", bank.Balance()) // A2
	}()
	n.Add(1)
	go func() {
		defer n.Done()
		bank.Deposit(100) // В
	}()

	n.Wait()

	fmt.Println("=", bank.Balance()) // C
}
