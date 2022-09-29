package main

/*Ejercicio 4 - Envios

Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de productos:
La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto requiere un
adicional al momento de ser enviado:

	Chico: Ningún adicional.
	Mediano: Requiere un %5 más de espacio
	Grande: Requiere un %20 más de espacio
	Frágil: Requiere un %75 más de espacio
	Especial: Sólo puede ser enviado con productos especiales

Para ello requerimos que los productos guarden el tamaño y tengan un método Tamaño Total que nos
devuelva el espacio en cm3 que requerimos para ser enviado.

Y una estructura Flete que tenga los métodos:
	1. Agregar Producto: agregar producto al flete
	2. Calcular Envios: calcula la cantidad de envíos que debe realizar sabiendo que solo puede
	cargar un total de 10.000.000 cm3 por envío.
*/

const (
	Chico    = "chico"
	Mediano  = "mediano"
	Grande   = "grande"
	Fragil   = "frágil"
	Especial = "especial"
)

type prod struct {
	tamanio    int64
	tamanioCm3 float64
}

type flete struct {
	productos prod
}

func main() {

}
