/*Ejercicio 1 - Datos de clientes
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones.
Para ello, cuentan con todo el detalle necesario en un archivo .txt.

Es necesario desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo,
no han pasado el archivo a leer por nuestro programa.
Desarrolla el código necesario para leer los datos del archivo llamado “customers.txt” (recuerda lo visto sobre el pkg “os”).

Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso,
el programa deberá arrojar un panic al intentar leer un archivo que no existe,
mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func crearArchivo() {
	data := "info"

	err := os.WriteFile("customers.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
	}

}

func lectura() {
	file, err := os.ReadFile("customers.txt")
	info := string(file)
	fmt.Println(strings.ReplaceAll(info, ";", "\t\t\t"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("el archivo contiene %s \n", file)
}

func main() {
	fmt.Println("inciando...")

	defer func() {
		fmt.Println("ejecucion finalizada...")
	}()

	crearArchivo()
	lectura()

}
