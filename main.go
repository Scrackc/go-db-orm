package main

import (
	"fmt"

	"github.com/Scrackc/go-db-orm/model"
	"github.com/Scrackc/go-db-orm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	// Migrar tablas
	storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})

	// product := model.Product{}
	// product.ID = 2

	// storage.DB().Unscoped().Delete(&product)
	// invoice := model.InvoiceHeader{
	// 	Client: "Eduardo Pérez",
	// 	InvoiceItems: []model.InvoiceItem{
	// 		{ProductId: 1},
	// 		{ProductId: 2},
	// 	},
	// }
	// storage.DB().Create(&invoice)
	var name string
	fmt.Scan(&name)

}
