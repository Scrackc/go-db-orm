package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver de storage
type Driver string

// Drivers
const (
	MySQl    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// Crea la conxión con la base de datos
func New(d Driver) {
	switch d {
	case MySQl:
		newMySqlDB()
	case Postgres:
		newPGDB()
	}
}

func newPGDB() {
	// Singleton solo se va a ejeuctar una vez aun que se llame muchas veces
	const dsn = "host=localhost user=postgres password=postgresPassword dbname=go-db port=5432 sslmode=disable"
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			// panic(err)
			log.Fatalf("No se puede conectar a la base de datos: %v", err)
		}
		fmt.Println("Conectado a PG")
	})
}
func newMySqlDB() {
	// Singleton solo se va a ejeuctar una vez aun que se llame muchas veces
	const dsn = "root:postgresPassword@tcp(127.0.0.1:3306)/go-db?charset=utf8mb4&parseTime=True&loc=Local"
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			// panic(err)
			log.Fatalf("No se puede conectar a la base de datos: %v", err)
		}
		// defer db.Close()
		fmt.Println("Conectado a MYSQL")
	})
}

// DB retorna una unica instancia de DB
func DB() *gorm.DB {
	return db
}
