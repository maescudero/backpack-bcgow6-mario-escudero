package main

import (
	"errors"
	"fmt"
)

func calcProm(notas ...float64) float64 {
	contador := 0.0
	for _, value := range notas {
		if value < 0 {
			errors.New("ingrese valor igual mayor a 0")
		} else {
			contador += value
		}
	}

	promedio := contador / float64(len(notas))

	return promedio
}

func main() {
	x := calcProm(10, 8.6, 4.6)

	fmt.Println(x)
}
