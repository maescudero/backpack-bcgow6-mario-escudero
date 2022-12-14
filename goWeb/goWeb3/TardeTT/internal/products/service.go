package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, productType string, count int, price float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateNameAndType(id int, nombre string, tipo string) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll() //llamamos al repository y esperamos su reespuesya
	if err != nil {
		return nil, err //procesamos
	}

	return ps, nil // lo va a recibir el handler
}

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, tipo, cantidad, precio)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {

	return s.repository.Update(id, name, productType, count, price)
}

func (s *service) UpdateNameAndType(id int, name, productType string) (Product, error) {

	return s.repository.UpdateNameAndType(id, name, productType)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
