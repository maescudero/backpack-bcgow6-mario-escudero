package operaciones

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	//Arrange
	num1 := 10
	num2 := 5
	resExpected := 5
	errExpected := fmt.Sprintf("%d no puede ser menor a %d", num1, num2)
	//Act
	res := Restar(num1, num2)
	assert.Equal(t, resExpected, res, errExpected)
}

func TestOrdenar(t *testing.T) {
	//Arrange
	numeros := [...]int{18, 28, 21, 96, 97}

	resExpected := 97

	errEx := fmt.Sprintf("%d debe ser el mayor", resExpected)
	//Act
	res := Order(numeros[:])

	assert.Equal(t, resExpected, res, errEx)

}

func TestDividir(t *testing.T) {
	//Arrange
	numerador := 10
	denominador := 5
	resExpected := 2
	errExpected := fmt.Sprintf("denominador tiene que ser mayor a 0 -> %d", denominador)
	//Act
	res, err := Dividir(numerador, denominador)

	//Assert
	assert.Equal(t, resExpected, res, errExpected)
	assert.Nil(t, err)

}
