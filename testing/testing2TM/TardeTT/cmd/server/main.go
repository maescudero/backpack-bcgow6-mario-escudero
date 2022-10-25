package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/cmd/server/handler"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/docs"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/internal/products"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/pkg/store"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Bootcamp Go Wave 6 - API
// @version         1.0
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   API Support Digital House
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

func main() {

	_ = godotenv.Load()

	db := store.NewFileStore("./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.DELETE("/:id", p.Delete())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNameAndType())

	err := r.Run()

	if err != nil {
		return
	}
}
