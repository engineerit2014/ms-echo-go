package dto

import (
	"github.com/go-playground/validator/v10"
)

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

func (uu *UpdateUserRequest) Validate() error {
	if err := validator.New().Struct(uu); err != nil {
		return err
	}

	return nil
}
