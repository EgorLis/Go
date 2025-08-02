package tempconv

type Celsius float64

type Fahranheit float64

var Smt int = 1

const (
	AbsolutZeroC Celsius = -273.15
	BoilingC     Celsius = 100
	FreezingC    Celsius = 0
)

// При иницилизации пакета выполняется эта функция
func init() { Smt = 100 }

// Закомментил мейн, чтоб можно было проверить экспорт пакета
/*
func main() {
	c := Celsius(10)
	b := Fahranheit(100)

	// c = b - данное выражение ошибочно, это разные типы, хоть и основаны на одном и том же float64

	fmt.Printf("c %.2f = f %v, f %v = c %.2f \n", c, CToF(c), b, FToC(b))

	fmt.Println(c.String())
}*/
