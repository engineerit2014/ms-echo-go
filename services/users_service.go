package services

import (
	"context"
	"fmt"
	respKit "github.com/laironacosta/kit-go/middleware/responses"
	"github.com/laironacosta/ms-echo-go/controllers/dto"
	"github.com/laironacosta/ms-echo-go/enums"
	repo "github.com/laironacosta/ms-echo-go/repository"
	"strings"
)

type UserServiceInterface interface {
	Create(ctx context.Context, request dto.CreateUserRequest) error
	GetByEmail(ctx context.Context, email string) (*dto.User, error)
	UpdateByEmail(ctx context.Context, request dto.UpdateUserRequest, email string) error
	DeleteByEmail(ctx context.Context, email string) error
}

type UserService struct {
	userRepo repo.UserRepositoryInterface
}

func NewUserService(userRepo repo.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepo,
	}
}

func (s *UserService) Create(ctx context.Context, request dto.CreateUserRequest) error {
	if err := s.userRepo.Create(ctx, request); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	fmt.Printf("Service email received: %+v \n", email)
	email = strings.TrimSpace(email)
	fmt.Printf("Service email received2: %+v \n", email)
	if email == "" {
		return &dto.User{}, respKit.GenericBadRequestError(enums.ErrorEmailNotEmptyCode, enums.ErrorEmailNotEmptyMsg)
	}

	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return &dto.User{}, err
	}

	return u, nil
}

func (s *UserService) UpdateByEmail(ctx context.Context, request dto.UpdateUserRequest, email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return respKit.GenericBadRequestError(enums.ErrorEmailNotEmptyCode, enums.ErrorEmailNotEmptyMsg)
	}

	if err := s.userRepo.UpdateByEmail(ctx, request, email); err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteByEmail(ctx context.Context, email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return respKit.GenericBadRequestError(enums.ErrorEmailNotEmptyCode, enums.ErrorEmailNotEmptyMsg)
	}

	err := s.userRepo.DeleteByEmail(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
