package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationUpdate(t *testing.T) {
	//Arrange.
	expectedDB := Product{

		ID:    1,
		Name:  "pepsi",
		Type:  "gaseosa",
		Count: 12,
		Price: 9,
	}

	initialDB := []Product{
		{
			ID:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Type:  "Galletitas y snacks",
			Count: 2000,
			Price: 300,
		},
	}

	mockStorage := MockStorage{
		dataMock: initialDB,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	//Act.
	productToUpdate := Product{
		Name:  "pepsi",
		Type:  "gaseosa",
		Count: 12,
		Price: 9,
	}

	result, err := service.Update(1, productToUpdate.Name, productToUpdate.Type, productToUpdate.Count, productToUpdate.Price)

	//Assert.
	assert.Nil(t, err)
	assert.Equal(t, expectedDB, result)
	//assert.Equal(t, productToUpdate, result)

}

func TestServiceIntegrationDelete(t *testing.T) {
	//Arrange
	expectedDB := []Product{
		{
			ID:    2,
			Name:  "Raqueta",
			Type:  "Tenis",
			Count: 15,
			Price: 20000,
		},
	}

	initialDB := []Product{
		{
			ID:    1,
			Name:  "Caja de galletitas Boreo 1kg",
			Type:  "Galletitas y snacks",
			Count: 2000,
			Price: 300,
		}, {
			ID:    2,
			Name:  "Raqueta",
			Type:  "Tenis",
			Count: 15,
			Price: 20000,
		},
	}
	mockStorage := MockStorage{
		dataMock: initialDB,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	//Act

	result := service.Delete(1)

	//Assert
	assert.Nil(t, result)
	assert.Equal(t, expectedDB, mockStorage.dataMock)
}
