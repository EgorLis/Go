// Lissajous генерирует анимированный GIF из
// случайных фигур Лиссажу
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{R: 0, G: 255, B: 0, A: 255}}

const (
	blackIndex = 0 // Первый цвет палитры
	greenIndex = 1 // Следующий цвет палитры
)

func main() {
	f, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	defer f.Close()
	// рисуем в этот файл
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // количество полных колебаний x
		res     = 0.001 // угловое разрешение
		size    = 100   // канва изображения охватывает [size..+size]
		nframes = 64    // количество кадров анимации
		delay   = 8     // задержка между кадрами (единица - 10 мс)
	)
	rand.New(rand.NewSource((time.Now().UTC().UnixNano())))
	freq := rand.Float64() * 3.0 // относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim) // примечание: игнорируем ошибки
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
}
