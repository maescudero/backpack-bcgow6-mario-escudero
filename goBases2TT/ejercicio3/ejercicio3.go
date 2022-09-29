package main

import "strings"

type tienda struct {
	productos producto
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type Producto interface {
	calcularCosto()
}

func (p producto) calcularCosto(tipo string, precio float64) float64 {
	if strings.ToLower(tipo) == "mediano" {
		p.precio += precio * 0.03
	} else if strings.ToLower(tipo) == "grande" {
		p.precio = precio*0.06 + 2500
	} else {
		return precio
	}
	return precio
}

type Ecommerce interface {
	total()
	agregar()
}

func (p producto) nuevoProducto(tipo string, nombre string, precio float64) (string, string, float64) {
	p.tipo = tipo
	p.nombre = nombre
	p.precio = precio

	return tipo, nombre, precio
}

func nuevaTienda(tipo string) Ecommerce {
	return nil
}

func main() {

}
