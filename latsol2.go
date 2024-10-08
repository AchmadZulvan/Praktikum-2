package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari float64
	fmt.Print("Masukkan jejari : ")
	fmt.Scan(&jejari)

	volumebola := (4 * math.Pi * jejari * jejari * jejari) / 3

	luasbola := 4 * math.Pi * math.Pow(jejari, 2)

	fmt.Print(" Bola dengan jejari ", jejari, " memiliki volume ", volumebola, " dan luas kulit ", luasbola)

}
