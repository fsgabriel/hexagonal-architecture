package application

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
)

const (
	DISABLE = "disable"
	ENABLE  = "enable"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

var (
	ErrInvalidPrice  = errors.New("the price must be greater than zero to enable the product")
	ErrDisable       = errors.New("the price must be zero in order to have the product disable")
	ErrInvalidStatus = errors.New("status must be enable or disable")
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLE
	}

	if p.Status != DISABLE && p.Status != ENABLE {
		return false, ErrInvalidStatus
	}

	if p.Price < 0 {
		return false, ErrInvalidPrice
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, fmt.Errorf("%w: missing required fields", err)
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLE
		return nil
	}
	return ErrInvalidPrice
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLE
		return nil
	}
	return ErrDisable
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
