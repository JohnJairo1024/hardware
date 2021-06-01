package models

import "errors"

type ProductStatus uint8

var (
	ErrProductEmptyName = errors.New("product can't be empty")
)

const (
	ProductStatus_Unavailable = 0
	productStatus_Available   = 1
)

type Product struct {
	Model
	Name     string        `gorm:"size:512;not null;unique" json:"name"`
	Price    float64       `gorm:"type:decimal(10,2);not null;default:0.0" json:"price"`
	Quantity uint16        `gorm:"default:0;unsigned" json:"quantity"`
	Status   ProductStatus `gorm:"char(1);default:0" json:"status"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrProductEmptyName
	}
	return nil
}
