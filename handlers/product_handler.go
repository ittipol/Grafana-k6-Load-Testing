package handlers

import (
	"go-load-testing/services"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandler {
	return &productHandler{productService}
}

func (obj productHandler) GetProducts(c *fiber.Ctx) error {

	products, err := obj.productService.GetProducts()

	if err != nil {
		c.Status(fiber.StatusBadGateway)
		return c.JSON("Error")
	}

	response := fiber.Map{
		"status":   fiber.StatusOK,
		"products": products,
	}

	c.Status(fiber.StatusOK)
	return c.JSON(response)
}
