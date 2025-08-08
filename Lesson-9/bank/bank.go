// Пакет bank предоставляет безопасный с точки зрения
// параллельности банк с одним счетом.
package bank

var deposits = make(chan int) // отправление вклада
var balances = make(chan int) // получение баланса

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance ограничен go-подпрограммой teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
