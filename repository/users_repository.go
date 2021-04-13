package repository

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/laironacosta/ms-echo-go/controllers/dto"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, request dto.CreateUserRequest) error
	GetByEmail(ctx context.Context, email string) (*dto.User, error)
	UpdateByEmail(ctx context.Context, request dto.UpdateUserRequest, email string) error
	DeleteByEmail(ctx context.Context, email string) error
}

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) UserRepositoryInterface {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Create(ctx context.Context, request dto.CreateUserRequest) error {
	u := dto.User{
		Name:  request.Name,
		Email: request.Email,
	}
	_, err := r.db.Model(&u).Context(ctx).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	u := dto.User{}
	err := r.db.Model(&u).Context(ctx).Where("email = ?", email).Select()
	if err != nil {
		return &u, err
	}
	return &u, nil
}

func (r *UserRepository) UpdateByEmail(ctx context.Context, request dto.UpdateUserRequest, email string) error {
	u := dto.User{}
	_, err := r.db.Model(&u).Context(ctx).Set("name = ?", request.Name).Where("email = ?", email).Update()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteByEmail(ctx context.Context, email string) error {
	u := dto.User{}
	_, err := r.db.Model(&u).Context(ctx).Where("email = ?", email).Delete()
	if err != nil {
		return err
	}
	return nil
}
