package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Ejercicio 1 - Crear Entidad
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).

Ejercicio 2 - Validación de campos
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo).

Ejercicio 3 - Validar Token
Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::

1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”.


*/

type Product struct {
	Id     int
	Nombre string  `json:"nombre" binding:"required"`
	Color  string  `json:"color" binding:"required"`
	Precio float64 `json:"precio" binding:"required"`
	Stock  int64   `json:"stock" binding:"required"`
}

var productos []Product
var Id int

func main() {
	router := gin.Default()

	router.POST("/registro", registo())

	router.Run()

}

func registo() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token != "messi" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
			return
		}

		var req Product
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		comprobar := req.check(ctx)

		if comprobar {
			Id++
			req.Id = Id
			productos = append(productos, req)
			ctx.JSON(200, productos)
		}

	}

}

func (p Product) check(ctx *gin.Context) bool {
	if p.Nombre == "" {
		respuesta := fmt.Sprintf("el ingreso %s no es valido", p.Nombre)
		ctx.JSON(http.StatusBadRequest, respuesta)
		return false
	}
	if p.Color == "" {
		respuesta := fmt.Sprintf("el ingreso %s no es valido", p.Color)
		ctx.JSON(http.StatusBadRequest, respuesta)
		return false
	}
	if p.Precio < 0 {
		respuesta := fmt.Sprintf("el ingreso %f no es valido", p.Precio)
		ctx.JSON(http.StatusBadRequest, respuesta)
		return false
	}
	if p.Stock < 0 {
		respuesta := fmt.Sprintf("el ingreso %v no es valido", p.Stock)
		ctx.JSON(http.StatusBadRequest, respuesta)
		return false
	}

	return true

}
