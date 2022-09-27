package main

type usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos *[]producto
}

type producto struct {
	Nombre   string
	Apellido string
	Cantidad int
}

func (p *producto) nuevoProducto(nombre string, apellido string, cantidad int) {
	p.Nombre = nombre
	p.Apellido = apellido
	p.Cantidad = cantidad
}

func main() {

}
