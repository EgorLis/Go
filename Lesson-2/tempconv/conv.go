package tempconv

import "fmt"

func (c Celsius) String() string { return fmt.Sprintf("%.1f°С", c) }

func CToF(c Celsius) Fahranheit {
	return Fahranheit(c*9/5 + 32)
}

func FToC(f Fahranheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
