# CRUD GO CON GORM (ORM)

1. Levantar db 
```
docker compose up -d
```
2. Ejecutar 
```
go run main.go
```

#### CONEXIÓN CON LA BASE DE DATOS
```
    // Motor de base de datos a conectar 
    driver := storage.Postgres
	storage.New(driver)
```

## Migrar tablas
```
    storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})
```
## Crear un registro
```
    // La observacion es opcional
    obs := "testing with go"
	product := model.Product{
		Name:         "Curso de Testing",
		Price:        500,
		Observations: &obs,
	}
    torage.DB().Create(&product)
```
## Crear multiples registros (lote)
```
    product1 := model.Product{
		Name:  "Curso de GO",
		Price: 300,
	}
	obs := "testing with go"
	product2 := model.Product{
		Name:         "Curso de Testing",
		Price:        500,
		Observations: &obs,
	}
	product3 := model.Product{
		Name:  "Curso de JS",
		Price: 200,
	}
	products := []model.Product{product1, product2, product3}
	storage.DB().Create(&products)
```
## Obtener todos los registros
```
    products := make([]model.Product, 0)
	storage.DB().Find(&products)
	for _, product := range products {
		fmt.Printf("%d - %s \n", product.ID, product.Name)
	}
```
## Obtener un registro
```
    product := model.Product{}
	storage.DB().First(&product, 2)
	fmt.Printf("%d - %s\n", product.ID, product.Name)
```
## Actualizar registro
```
    product := model.Product{}
    product.ID = 2
	storage.DB().Model(&product).Updates(model.Product{Name: "Curso Testing", Price: 150})
```
## Eliminar un registro (Eliminación suave)
```
    product := model.Product{}
	product.ID = 2
	storage.DB().Delete(&product)
```
## Eliminar un registro (Permanente) 
```
    product := model.Product{}
	product.ID = 2
	storage.DB().Unscoped().Delete(&product)
```
## Transacción crear factura
```
	invoice := model.InvoiceHeader{
		Client: "Eduardo Pérez",
		InvoiceItems: []model.InvoiceItem{
			{ProductId: 1},
			{ProductId: 2},
		},
	}
	storage.DB().Create(&invoice)
```