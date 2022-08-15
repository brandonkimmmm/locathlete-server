package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidationUtil struct {
	validator *validator.Validate
}

func NewValidationUtil() echo.Validator {
	return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}