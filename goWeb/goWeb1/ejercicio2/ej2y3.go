package main

/*
Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.
*/

import (
	//"net/http"

	//"fmt"
	//	"os"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb1/ejercicio1"
)

/*func get() string {

	data, _ := os.ReadFile("/goWeb/goWeb1/ejercicio1/productos.json")

	fmt.Println(string(data))

	return string(data)
} */

func main() {

	router := gin.Default()

	data, err := os.ReadFile("./productos.json")

	if err != nil {
		fmt.Println("omar algo anda mal")
	}
	x := string(data)

	router.GET("/get", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"data": x,
		})
	})

	router.GET("/saludo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Mario :v",
		})
	})

	router.Run()

}
