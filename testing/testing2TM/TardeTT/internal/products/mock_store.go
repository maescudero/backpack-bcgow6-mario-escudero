package products

import "github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/internal/domains"

type MockStorage struct {
	dataMock   []domains.Product
	errOnWrite error
	errOnRead  error
}

func (m *MockStorage) Read(data interface{}) (err error) {
	if m.errOnRead != nil {
		return m.errOnRead
	}

	castedData := data.(*[]domains.Product)
	*castedData = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}

	castedData := data.(*domains.Product)
	m.dataMock = append(m.dataMock, *castedData)
	return nil
}
