//Ejercicio 1 - Red social
//Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura.
//Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa
//y para las funciones.
//La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
//Y deben implementarse las funciones:
//Cambiar nombre: me permite cambiar el nombre y apellido.
//Cambiar edad: me permite cambiar la edad.
//Cambiar correo: me permite cambiar el correo.
//Cambiar contraseña: me permite cambiar la contraseña.

package main

import "fmt"

type persona struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (p *persona) cambiarNombre(nombre string, apellido string) {
	p.Nombre = nombre
	p.Apellido = apellido
}

func (p *persona) cambiarEdad(edad int) {
	p.Edad = edad
}

func (p *persona) cambiarCorreo(correo string) {
	p.Correo = correo
}

func (p *persona) cambiarContraseña(contraseña string) {
	p.Contraseña = contraseña
}

func main() {

	p := persona{
		Nombre:     "Marrito",
		Apellido:   "Escudero",
		Edad:       23,
		Correo:     "mario@mercadolibre.com",
		Contraseña: "hola123F",
	}
	fmt.Println("antes", p)
	p.cambiarContraseña("Hola123")
	p.cambiarNombre("pablo", "juarez")

	fmt.Println("despues:", p)

}
