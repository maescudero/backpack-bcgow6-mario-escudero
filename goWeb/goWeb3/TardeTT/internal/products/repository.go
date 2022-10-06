package products

import (
	"fmt"

	"github.com/maescudero/backpack-bcgow6-mario-escudero/goWeb/goWeb3/TardeTT/pkg/store"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product
var lastID int

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

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) UpdateNameAndType(id int, nombre string, tipo string) (Product, error) {
	var updated bool
	p := Product{Name: nombre, Type: tipo}
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = nombre
			ps[i].Type = tipo
			p = ps[i]
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var deleted bool
	var pos int
	for i := range ps {
		if ps[i].ID == id {
			pos = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("product id %d not exists", id)
	}

	ps = append(ps[:pos], ps[pos+1:]...)
	return nil
}
