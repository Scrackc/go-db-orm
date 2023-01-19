package model

import (
	"gorm.io/gorm"
)

// Product model of table products
type Product struct {
	gorm.Model
	Name         string        `gorm:"type:varchar(100); not null"`
	Observations *string       `gorm:"type:varchar(100)"`
	Price        int           `gorm:"not null"`
	InvoiceItems []InvoiceItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ProductId;references:id"`
}

// InvoiceHeader model of table  invoice_headers
type InvoiceHeader struct {
	gorm.Model
	Client       string        `gorm:"type:varchar(100); not null"`
	InvoiceItems []InvoiceItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:InvoiceHeaderId;references:id"`
}

// InvoiceItem model of table invoice_items
type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderId uint
	ProductId       uint
}
