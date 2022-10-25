package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/internal/domains"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/internal/products"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/pkg/web"
)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(401, web.NewResponse(500, nil, err.Error()))
			return
		}

		if len(p) == 0 {
			ctx.JSON(401, web.NewResponse(404, nil, "no hay productos"))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*	token := ctx.Request.Header.Get("token")
			if token != os.Getenv("TOKEN") {
				ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
				return
			} */

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, web.NewResponse(401, nil, err.Error()))
			return
		}

		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(401, nil, "el nombre del prod es requerido"))
			return
		}
		if req.Type == "" {
			ctx.JSON(400, web.NewResponse(401, nil, "el tipo del prod es requerido"))
			return
		}
		if req.Count == 0 {
			ctx.JSON(400, web.NewResponse(401, nil, "la cantidad del prod es requerido"))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(401, nil, "el precio del prod es requerido"))
			return
		}

		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(400, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*	token := ctx.GetHeader("token")
			if token != os.Getenv("TOKEN") {
				ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
				return
			} */

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "No fue posible obtener el id del producto"),
			)
			return
		}

		var request domains.Product

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		//	errorParam := validateRequest(request)

		/*if errorParam != "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, errorParam))
			return
		} */

		product, err := c.service.Update(idProduct, request.Nombre, request.Tipo, request.Cantidad, request.Precio)
		fmt.Println(product)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))
	}
}

/*if req.Name == "" {
	ctx.JSON(400, web.NewResponse(401, nil, "el nombre del prod es requerido"))
	return
}
if req.Type == "" {
	ctx.JSON(400, web.NewResponse(401, nil, "el tipo del prod es requerido"))
	return
}
if req.Count == 0 {
	ctx.JSON(400, web.NewResponse(401, nil, "la cantidad del prod es requerido"))
	return
}
if req.Price == 0 {
	ctx.JSON(400, web.NewResponse(401, nil, "el precio del prod es requerido"))
	return
} */

func (p *Product) UpdateNameAndType() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(401, web.NewResponse(400, nil, "ID invalido"))
			//ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.UpdateNameAndType(int(id), req.Name, req.Type)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, web.NewResponse(400, nil, "ID invalido"))
			//ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		deleted := p.service.Delete(id)

		if deleted != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, deleted.Error()))
			return
		}

		c.JSON(http.StatusNoContent, "Deleted")

	}
}
