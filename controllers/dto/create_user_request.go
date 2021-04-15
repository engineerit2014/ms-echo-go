package dto

import (
	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (cu *CreateUserRequest) Validate() error {
	if err := validator.New().Struct(cu); err != nil {
		return err
	}

	return nil
}
