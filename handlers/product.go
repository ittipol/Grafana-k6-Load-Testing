package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Job struct {
	Type   string `validate:"required,min=3,max=32"`
	Salary int    `validate:"required,number"`
}

type User struct {
	Name string `validate:"required,min=3,max=32"`
	// use `*bool` here otherwise the validation will fail for `false` values
	// Ref: https://github.com/go-playground/validator/issues/319#issuecomment-339222389
	IsActive *bool  `validate:"required"`
	Email    string `validate:"required,email,min=6,max=32"`
	Job      Job    `validate:"dive"`
}

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
}

// func validate() {
// 	var validate = validator.New()
// }
