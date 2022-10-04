package main

import "fmt"

func impuesto(salario float64) float64 {
	impuestoMenor := 0.0
	impuestoMayor := 0.0

	/*if salario < 0 {
		err = errors.New("no puede ser negativo")
		return
	} */
	if salario > 50000 {
		impuestoMenor = salario * 0.17
		if salario > 150000 {
			impuestoMayor = salario * 0.10
		}
	}

	totalTax := impuestoMenor + impuestoMayor

	return totalTax
}

func main() {
	fmt.Println(impuesto(200000))
}
