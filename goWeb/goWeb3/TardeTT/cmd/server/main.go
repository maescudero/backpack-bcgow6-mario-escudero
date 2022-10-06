package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb3/TardeTT/cmd/server/handler"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb3/TardeTT/internal/pkg/store"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb3/TardeTT/internal/products"
)

func main() {

	_ = godotenv.Load()

	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	r.Run()
}
