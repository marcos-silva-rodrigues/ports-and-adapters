package main

import (
	"database/sql"

	dbPackage "github.com/marcos-silva-rodrigues/go-hexagonal/adapters/db"
	"github.com/marcos-silva-rodrigues/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := dbPackage.NewProductDb(db)
	produceService := application.NewProductService(productDbAdapter)
	product, _ := produceService.Create("Product Example", 30)

	produceService.Enable(product)
}
