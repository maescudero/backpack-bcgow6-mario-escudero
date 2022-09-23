package main

import "fmt"

func main() {
	edad := 23
	var empleado bool = true
	antiguedad := 0.8
	sueldo := 20000

	if edad >= 22 && empleado == true && antiguedad >= 1 {
		fmt.Println("tiene acceso a prestamo")
		if sueldo > 100000 {
			fmt.Printf("y su prestamo es sin interes")
		}

	} else {
		fmt.Printf("no puede acceder a prestamos")
	}

}
