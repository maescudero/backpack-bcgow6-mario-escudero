package products

import (
	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb3/Mañana/internal/domains"
)

type Service interface {
	GetAll() ([]domains.Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (domains.Product, error)
	Update(id int, nombre, tipo string, cantidad int, precio float64) (domains.Product, error)
	UpdateNameAndType(id int, nombre string, tipo string) (domains.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]domains.Product, error) {
	return s.repository.GetAll()
}

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (domains.Product, error) {
	return s.repository.Store(nombre, tipo, cantidad, precio)
}

func (s *service) Update(id int, nombre, tipo string, cantidad int, precio float64) (domains.Product, error) {
	return s.repository.Update(id, nombre, tipo, cantidad, precio)
}

func (s *service) UpdateNameAndType(id int, nombre string, tipo string) (domains.Product, error) {
	return s.repository.UpdateNameAndType(id, nombre, tipo)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
