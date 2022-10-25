package test

import (
	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/internal/products"
)

type MockStorage struct {
	dataMock   []products.Product
	errOnWrite error
	errOnRead  error
	called     bool
}

func (m *MockStorage) Read(data interface{}) (err error) {
	if m.errOnRead != nil {
		return m.errOnRead
	}

	castedData := data.(*[]products.Product)
	*castedData = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}

	castedData := data.([]products.Product)
	m.dataMock = castedData
	//m.dataMock = append(m.dataMock, *castedData)
	return nil
}
