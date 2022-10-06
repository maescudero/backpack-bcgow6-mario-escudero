package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb3/Mañana/cmd/server/handler"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb3/Mañana/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNameAndType())
	pr.DELETE("/:id", p.Delete)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
