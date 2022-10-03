/*
Ejercicio 1 - Estructura un JSON
Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no),
fecha de creación.
Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string),
 fecha de transacción.
Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID        string
	Nombre    string
	Color     string
	Precio    float64
	Stock     int64
	Codigo    string
	Publicado bool
	fecha     string
}

func main() {

	p1 := Product{
		ID:        "123",
		Nombre:    "mouse",
		Color:     "negro",
		Precio:    70.5,
		Stock:     10,
		Codigo:    "HH100",
		Publicado: true,
		fecha:     "09/12/2018",
	}

	jsonData, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("mal..")
	}
	escritura := os.WriteFile("productos.json", []byte(jsonData), 0644)
	fmt.Println(escritura)
	fmt.Printf("producto: \n %s ", string(jsonData))
}
