// mandelbrot создает PNG - изображение фрактала Мандельброта
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 4096, 4096
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// точка (px, py) представляет комплексное значение z
			img.Set(px, py, mandelbrot(z))
		}
	}

	f, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	defer f.Close()
	// рисуем в этот файл

	png.Encode(f, img) // примечание: игнорируем ошибки
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// нормируем шаг n в диапазон [0,1]
			hue := float64(n) / iterations
			// насыщенность=1, яркость=1
			return hsvToRGB(hue, 1, 1)
		}
	}
	return color.Black
}

// Конвертация HSV в RGBA
func hsvToRGB(h, s, v float64) color.Color {
	h = math.Mod(h*360, 360) / 60
	i := math.Floor(h)
	f := h - i
	p := v * (1 - s)
	q := v * (1 - s*f)
	t := v * (1 - s*(1-f))

	var r, g, b float64
	switch int(i) {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	default: // case 5:
		r, g, b = v, p, q
	}
	return color.RGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: 255,
	}
}
