package products

import (
	"fmt"

	"github.com/maescudero/backpack-bcgow6-mario-escudero/testing/testing2TM/TardeTT/pkg/store"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product

//var lastID int

// ***Importado**//
type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, productType string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateNameAndType(id int, nombre string, tipo string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
} //struct implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {

	var ps []Product

	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}

	p := Product{id, nombre, tipo, cantidad, precio}

	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *repository) GetAll() (products []Product, err error) {

	err = r.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	err := r.db.Read(ps)
	if err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (result Product, err error) {
	var products []Product
	r.db.Read(&products)

	p, i := findEntityByID(id, products)
	if p.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
		return
	}

	if p.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
		return
	}

	result = Product{ID: id, Name: name, Type: productType, Count: count, Price: price}

	products[i] = result

	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}

	return
}

func findEntityByID(id int, products []Product) (p Product, index int) {

	for i, pr := range products {
		if pr.ID == id {
			p = pr
			index = i
			break
		}
	}

	return
}

func (r *repository) UpdateNameAndType(id int, nombre string, tipo string) (result Product, err error) {
	var products []Product
	r.db.Read(&products)

	p, i := findEntityByID(id, products)

	if p.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
		return
	}

	p.Name = nombre
	p.Type = tipo

	products[i] = p

	result = p

	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}

	return
}

func (r *repository) Delete(id int) (err error) {
	var products []Product
	r.db.Read(&products)

	p, i := findEntityByID(id, products)

	if p.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
		return
	}
	products = append(products[:i], products[i+1:]...)

	if err := r.db.Write(products); err != nil {
		return err
	}

	return
}
