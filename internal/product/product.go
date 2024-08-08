package product

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	IsActive    bool
}

func (p *Product) Activate() {
	p.IsActive = true
}

func (p *Product) Deactivate() {
	p.IsActive = false
}

func (p *Product) SetPrice(price float64) {
	p.Price = price
}

func NewProduct(name, description string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		IsActive:    true,
	}
}

type manager struct {
	repo Repository
}

type Repository interface {
	Save(product Product) error
	Get(id string) (Product, error)
	Delete(id string) error
}

func NewManager(repo Repository) *manager {
	return &manager{
		repo: repo,
	}
}

func (m *manager) SetPrice(id string, price float64) error {
	p, err := m.repo.Get(id)
	if err != nil {
		return err
	}
	p.SetPrice(price)
	return m.repo.Save(p)
}

func (m *manager) Save(product *Product) error {
	return m.repo.Save(*product)
}

func (m *manager) Activate(id string) error {
	p, err := m.repo.Get(id)
	if err != nil {
		return err
	}
	p.Activate()
	return m.repo.Save(p)
}

func (m *manager) Deactivate(id string) error {
	p, err := m.repo.Get(id)
	if err != nil {
		return err
	}
	p.Deactivate()
	return m.repo.Save(p)
}

func (m *manager) Get(id string) (Product, error) {
	return m.repo.Get(id)
}

func (m *manager) Delete(id string) error {
	return m.repo.Delete(id)
}

func (m *manager) New(name, description string, price float64) (Product, error) {
	if price <= 0 {
		return Product{}, ErrInvalidPrice
	}
	p := NewProduct(name, description, price)
	if err := m.repo.Save(*p); err != nil {
		return Product{}, err
	}
	return *p, nil
}

var ErrInvalidPrice = errors.New("price must be greater than 0")
