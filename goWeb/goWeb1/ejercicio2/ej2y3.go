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

	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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

func GetProducts() (products []Product, err error) {
	data, err0 := os.ReadFile("./productos.json")
	if err0 != nil {
		err = err0
	}

	err1 := json.Unmarshal(data, &products)
	if err1 != nil {
		err = err1
	}

	return
}

func main() {

	router := gin.Default()

	//creamos la ruta para devolver los productos
	router.GET("/products", GetAll)

	router.GET("/saludo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Mario :v",
		})
	})

	router.Run()

}

func GetAll(c *gin.Context) {

	var products []Product
	var err error

	products, err = GetProducts()

	if err != nil {
		c.JSON(500, err)
		return
	}

	var productsResponse []Product

	//logica de filtrado
	productName := c.Query("name")
	productColour := c.Query("colour")
	productPrice := c.Query("price")
	productStock := c.Query("stock")
	productCreatedAt := c.Query("createdAt")

	if productName == "" && productColour == "" && productPrice == "" && productStock == "" && productCreatedAt == "" {
		c.JSON(200, products)
	} else {

		priceKey, _ := strconv.ParseFloat(productPrice, 64)
		stockKey, _ := strconv.ParseInt(productStock, 0, 0)

		for _, product := range products {

			if (productName != "" && product.Nombre == productName) || (productColour != "" && product.Color == productColour) || (productPrice != "" && product.Precio == priceKey) || (productStock != "" && int(product.Stock) == int(stockKey)) || (productCreatedAt != "" && product.fecha == productCreatedAt) {
				productsResponse = append(productsResponse, product)
			}
		}

		c.JSON(200, productsResponse)

	}

}
