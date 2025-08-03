package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	defaultWidth, defaultHeight = 600, 320    // Размер канвы в пикселях
	cells                       = 100         // количество ячеек сетки
	xyrange                     = 30.0        // диапозон осей (-xyrange...+xyrange)
	angle                       = math.Pi / 6 // углы осей x, y (=30')
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", svg)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func svg(w http.ResponseWriter, r *http.Request) {
	// парсим query-параметры
	qs := r.URL.Query()

	// width
	wStr := qs.Get("width")
	width, err := strconv.Atoi(wStr)
	if err != nil || width <= 0 {
		width = defaultWidth
	}

	// height
	hStr := qs.Get("height")
	height, err := strconv.Atoi(hStr)
	if err != nil || height <= 0 {
		height = defaultHeight
	}

	// цвет обводки
	color := qs.Get("color")
	if color == "" {
		color = "grey"
	}

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:%s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j, width, height int) (float64, float64) {
	// ищем угловую точку (x, y) ячейки (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	xyscale := width / 2 / xyrange  // пикселей в единице x или y
	zscale := float64(height) * 0.4 // пикселей в единице z
	// вычисляем высоту поверхности z
	z := f(x, y)
	// изометрически проецируем (x, y, z) на двумерную канву SVG (sx, sy)
	sx := float64(width/2) + (x-y)*cos30*float64(xyscale)
	sy := float64(height/2) + (x+y)*sin30*float64(xyscale) - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // расстояние от (0, 0)
	return math.Sin(r) / r
}
