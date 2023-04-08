package main

import (
	"database/sql"
	adaptador "github.com/masilvasql/go-hexagonal/adapters/db"
	"github.com/masilvasql/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	defer db.Close()
	productDBAdapter := adaptador.NewProductDB(db)
	productService := application.NewProductService(productDBAdapter)
	product, err := productService.Create("Product 2", 10.0)
	if err != nil {
		panic(err)
	}
	product, err = productService.Get(product.GetId())

	if err != nil {
		panic(err)
	}
	product, err = productService.Enable(product)
	if err != nil {
		panic(err)
	}
	print(product.GetStatus())
}
