package operaciones

import "fmt"

func Restar(num1 int, num2 int) int {
	return num1 - num2
}

func Order(numeros []int) int {

	//numeros = [...]{18, 28, 21, 96, 97}
	numeroMayor := numeros[0]
	for _, numero := range numeros {
		if numero > numeroMayor {
			numeroMayor = numero
		}
	}
	//fmt.Printf("El número mayor entre %v es %d ", numeros, numeroMayor)

	return numeroMayor
}

func Dividir(num, den int) (int, error) {

	if den <= 0 {
		return 0, fmt.Errorf("a no puede ser:%d", den)
	}

	return num / den, nil
}
