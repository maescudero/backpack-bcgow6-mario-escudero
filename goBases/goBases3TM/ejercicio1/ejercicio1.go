package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	idProd := []int{1243, 12414, 100}
	precio := []float64{500.0, 12.0, 300.0}
	cantidad := []int{2, 4, 9}

	m := fmt.Sprintf("ID: %v  \t Precio: %v \t Cantidad: %v ", idProd, precio, cantidad)
	mensaje := []byte(m)
	fmt.Println("ID:  \t Precio:  \t Cantidad: ")
	for i := 0; i < len(idProd); i++ {

		m := fmt.Sprintf("%v  \t  %v \t\t %v  \n", idProd[i], precio[i], cantidad[i])
		io.WriteString(os.Stdout, string(m))
		err := os.WriteFile("./myFile.csv", mensaje, 0644)
		fmt.Println(err)
	}

	//io.WriteString(os.Stdout, string(x))

	//data ,err := os.ReadFile("./myFile.csv")

	//fmt.Println(data)

}
