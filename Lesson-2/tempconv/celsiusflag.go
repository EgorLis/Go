package tempconv

import (
	"flag"
	"fmt"
)

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // проверки ошибок не нужны
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahranheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("неверная температура %q", s)
}

// CelsiusFlag определяет флаг Celsius с указанным именем, значением
// по умолчанию и строкой-инструкцией по применению и возвращает адрес
// переменной-флага. Аргумент флага должен содержать числовое значение
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
