package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/cmd/server/handler"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/internal/products"
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/pkg/store"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("token", "123456")

	db := store.NewFileStore("./products.json")

	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")

	return req, httptest.NewRecorder()
}

/*
	func Test_SaveProduct_OK(t *testing.T) {
		r := createServer()
		req, rr := createRequestTest(http.MethodPost, "/products/",
			`{"nombre": "Televisor LCD",
	            "tipo": "electrodomesticos",
	            "cantidad": 5,
	            "precio": 20000}`)

		r.ServeHTTP(rr, req)
		assert.Equal(t, 200, rr.Code)
	}
*/
func Test_Update_OK(t *testing.T) {

	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/products/",
		`{"nombre": "Televisor LCD",
            "tipo": "electrodomesticos",
            "cantidad": 5,
            "precio": 20000}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	req, rr = createRequestTest(http.MethodPut, "/products/1",
		`{  "nombre": "Televisor LCD",
            "tipo": "electrodomesticos",
            "cantidad": 5,
            "precio": 20000}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

}
