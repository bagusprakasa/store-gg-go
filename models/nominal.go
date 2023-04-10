package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Nominal struct {
	gorm.Model
	Name     string
	Quantity decimal.Decimal `gorm:"type:double;"`
	Price    decimal.Decimal `gorm:"type:double;"`
}
