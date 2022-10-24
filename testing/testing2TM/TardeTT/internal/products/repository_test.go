package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByName(t *testing.T) {
	//arrange

	expectedData := []Product{
		{
			ID:    1,
			Name:  "joysticj",
			Type:  "HD",
			Count: 123,
			Price: 12321,
		},
		{
			ID:    2,
			Name:  "joy",
			Type:  "HDD",
			Count: 123123,
			Price: 123,
		},
	}

	stubStore := MockStorage{
		dataMock: expectedData,
	}

	repo := NewRepository(&stubStore)
	//act
	resultado, err := repo.GetAll()
	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedData, resultado)
}

func TestUpdateName(t *testing.T) {
	id, nombre, tipo := 1, "Update After", "nuevo"

	expectedProducts := Product{
		ID:    1,
		Name:  "Update After",
		Type:  "nuevo",
		Count: 123,
		Price: 12,
	}

	products := []Product{{
		ID:    1,
		Name:  "algo",
		Type:  "galleta",
		Count: 123,
		Price: 12,
	}}

	//data, _ := json.Marshal(products)
	mock := MockStorage{
		dataMock: products,
	}

	/*stubStore := store.FileStore{
		FileName: "",
		Mock:     &mock,
	} */

	r := NewRepository(&mock)
	productUpdated, err := r.UpdateNameAndType(id, nombre, tipo)
	assert.Nil(t, err)

	assert.Equal(t, expectedProducts, productUpdated, "el resultado es diferente a lo esperado")

}
