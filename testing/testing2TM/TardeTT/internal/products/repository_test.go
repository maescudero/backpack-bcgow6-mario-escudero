package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (s StubStore) Read(data interface{}) error {
	products := data.(*[]Product)
	stubData := []Product{
		{ID: 1,
			Name:  "joysticj",
			Type:  "HD",
			Count: 123,
			Price: 12321,
		},
		{ID: 2,
			Name:  "joy",
			Type:  "HDD",
			Count: 123123,
			Price: 123},
	}
	*products = stubData
	return nil
}

func (s StubStore) Write(data interface{}) error {
	return nil
}

func TestFindByName(t *testing.T) {
	//arrange
	myStubStore := StubStore{}
	repo := NewRepository(myStubStore)
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
}
