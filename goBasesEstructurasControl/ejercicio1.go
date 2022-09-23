package main

import "fmt"

func main() {
	palabra := "federer"
	//letras := []rune(palabra)
	/*for _, caracteres := range palabra {
		fmt.Printf(caracteres)
	} */

	fmt.Println("longitud: ", len(palabra))
	for i := 0; i < len(palabra); i++ {
		fmt.Println(palabra[i : i+1])
	}

}
