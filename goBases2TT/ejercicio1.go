package main

import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (p Persona) getPersona(nombre string, apellido string, dni int, fecha string) {

	fmt.Println("Nombre: ", p.Nombre, "Apellido:", p.Apellido, "DNI:", p.DNI, "fecha:", p.Fecha)
}

func main() {

	p := Persona{
		Nombre:   "Mario",
		Apellido: "Esc",
		DNI:      24124122542,
		Fecha:    "20/05/1999",
	}

	p.getPersona(p.Nombre, p.Apellido, p.DNI, p.Fecha)
}
