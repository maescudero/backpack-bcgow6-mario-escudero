package products

//"testing"

//"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/internal/domains"
//"github.com/stretchr/testify/assert"

type request struct {
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

/*func TestServiceUpdate(t *testing.T) {
	//Arrange.
	expectedDB := []domains.Product{
		{
			Id:       1,
			Nombre:   "Caja de galletitas Boreo 1kg",
			Tipo:     "Galletitas y snacks",
			Cantidad: 2000,
			Precio:   300,
		},
	}

	initialDB := []domains.Product{
		{
			Id:       1,
			Nombre:   "Caja de galletitas Boreo 1kg",
			Tipo:     "Galletitas y snacks",
			Cantidad: 2000,
			Precio:   300,
		},
	}

	mockStorage := MockStorage{
		dataMock: initialDB,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	//Act.
	productToUpdate := request{
		Nombre:   "pepsi",
		Tipo:     "gaseosa",
		Cantidad: 12,
		Precio:   9,
	}
	result, err := service.Update(1, productToUpdate.Nombre, productToUpdate.Tipo, productToUpdate.Cantidad, productToUpdate.Precio)

	//Assert.
	assert.Nil(t, err)
	assert.Equal(t, expectedDB, mockStorage.dataMock)
	assert.Equal(t, productToUpdate, result)

} */
