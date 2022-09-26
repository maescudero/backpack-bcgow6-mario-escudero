package main

import (
	"fmt"
)

type matrix struct {
	valores []int
	alto    int
	ancho   int
}

func (m matrix) setMatrix(alto int, ancho int, valores []int) {
	alto = m.alto
	ancho = m.ancho
	cuadrada := false
	if alto == ancho {
		cuadrada = true
	} else {
		cuadrada = false
	}

	/*for i := 0; i < alto * ancho; i++ {
		valores = append(valores, int(m.valores[i]))
		fmt.Println(valores[i])
	} */

	for i := 0; i < ancho*alto; i++ {

	}
	max := 0
	min := 0
	for _, value := range valores {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	fmt.Println("max:", max)

	fmt.Println(alto, ancho, "val: ", valores, cuadrada)

}

func main() {

	m := matrix{
		alto:    2,
		ancho:   2,
		valores: []int{2, 7, 4, 1},
	}

	m.setMatrix(m.alto, m.ancho, m.valores)
}
