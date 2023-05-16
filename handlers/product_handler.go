package handlers

import (
	"go-load-testing/services"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productService services.ProductService
	validator      *validator.Validate
}

func NewProductHandler(productService services.ProductService, validator *validator.Validate) ProductHandler {
	return &productHandler{productService, validator}
}

func (obj productHandler) GetProducts(c *fiber.Ctx) error {

	var user User

	err := obj.validator.Struct(user)

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
