package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb2/Tarde/cmd/server/handler"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb2/Tarde/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	r.Run()
}
