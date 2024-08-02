package product

import (
	"errors"

	"github.com/google/uuid"
)

type product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	IsActive    bool
}

func (p *product) Activate() {
	p.IsActive = true
}

func (p *product) Deactivate() {
	p.IsActive = false
}

func (p *product) SetPrice(price float64) {
	p.Price = price
}

func NewProduct(name, description string, price float64) *product {
	return &product{
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
	Save(product *product) error
	Get(id string) (*product, error)
	GetAll() ([]*product, error)
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

func (m *manager) Save(product *product) error {
	return m.repo.Save(product)
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

func (m *manager) Get(id string) (*product, error) {
	return m.repo.Get(id)
}

func (m *manager) GetAll() ([]*product, error) {
	return m.repo.GetAll()
}

func (m *manager) Delete(id string) error {
	return m.repo.Delete(id)
}

func (m *manager) New(name, description string, price float64) (*product, error) {
	if price <= 0 {
		return nil, ErrInvalidPrice
	}
	p := NewProduct(name, description, price)
	if err := m.repo.Save(p); err != nil {
		return nil, err
	}
	return p, nil
}

var ErrInvalidPrice = errors.New("price must be greater than 0")
