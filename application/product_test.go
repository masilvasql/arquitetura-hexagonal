package application_test

import (
	"github.com/masilvasql/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"

	require.Equal(t, "Product 1", product.GetName())
}

func TestProduct_GetId(t *testing.T) {
	product := application.Product{}
	product.ID = "123"

	require.Equal(t, "123", product.GetId())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED

	require.Equal(t, application.ENABLED, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	require.Equal(t, 10.0, product.GetPrice())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "the status must be ENBALED or DISABLED", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())

}
