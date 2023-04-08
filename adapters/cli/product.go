package cli

import (
	"fmt"
	"github.com/masilvasql/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {

	var result string

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s created with the name %s has been created with price %f and status %s",
			product.GetId(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been enabled with price %f and status %s",
			product.GetId(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been disabled with price %f and status %s",
			product.GetId(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	case "get":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been disabled with price %f and status %s",
			product.GetId(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())

	default:
		return result, fmt.Errorf("Action %s not found", action)
	}

	return result, nil
}
