package services

import (
	"context"
	"errors"
	"github.com/laironacosta/ms-echo-go/controllers/dto"
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
	e := strings.TrimSpace(email)
	if e == "" {
		return &dto.User{}, errors.New("email could not be empty")
	}

	u, err := s.userRepo.GetByEmail(ctx, e)
	if err != nil {
		return &dto.User{}, err
	}
	return u, nil
}

func (s *UserService) UpdateByEmail(ctx context.Context, request dto.UpdateUserRequest, email string) error {
	e := strings.TrimSpace(email)
	if e == "" {
		return errors.New("email could not be empty")
	}

	if err := s.userRepo.UpdateByEmail(ctx, request, e); err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteByEmail(ctx context.Context, email string) error {
	e := strings.TrimSpace(email)
	if e == "" {
		return errors.New("email could not be empty")
	}

	err := s.userRepo.DeleteByEmail(ctx, e)
	if err != nil {
		return err
	}
	return nil
}
