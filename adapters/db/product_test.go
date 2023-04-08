package db_test

import (
	"database/sql"
	"github.com/masilvasql/go-hexagonal/adapters/db"
	"github.com/masilvasql/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {

	table := `CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string
			);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES ("1", "Product 1", 10.00, "disabled");`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDB := db.NewProductDB(Db)
	product, err := productDB.Get("1")

	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.00, product.GetPrice())
	require.Equal(t, "1", product.GetId())
	require.Equal(t, "disabled", product.GetStatus())

}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDB(Db)

	product := application.NewProduct()
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Price = 10.00

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetId(), productResult.GetId())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetId(), productResult.GetId())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

}
