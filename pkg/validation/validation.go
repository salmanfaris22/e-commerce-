package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidateUser[T any](user T) error {
	err := validator.New().Struct(&user)
	if err != nil {
		return err
	}
	return nil
}
