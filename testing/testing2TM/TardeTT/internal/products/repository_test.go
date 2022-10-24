package products

import (
	"encoding/json"
	"testing"

	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func TestUpdateName(t *testing.T) {
	id, nombre, tipo := 1, "Update After", "nuevo"
	products := []*Product{{
		ID:    1,
		Name:  "algo",
		Type:  "galleta",
		Count: 123,
		Price: 12}}

	data, _ := json.Marshal(products)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	productUpdated, err := r.UpdateNameAndType(id, nombre, tipo)
	assert.Nil(t, err)

	assert.True(t, mock.ReadInvoked)
	assert.Equal(t, id, productUpdated.ID, productUpdated.Type)
	assert.Equal(t, nombre, productUpdated.Name, productUpdated.Type)
}

/*func TestFindByName(t *testing.T) {
	//arrange
	sut := StubStore{}
	r := NewRepository(&stubStore)
	dataEsperada := []Product{
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
			Price: 123},
	}
	//act
	resultado, _ := repo.GetAll()
	//assert
	assert.Equal(t, dataEsperada, resultado)
} */
