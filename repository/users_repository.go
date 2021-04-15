package repository

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	respKit "github.com/laironacosta/kit-go/middleware/responses"
	"github.com/laironacosta/ms-echo-go/controllers/dto"
	"github.com/laironacosta/ms-echo-go/enums"
	"strings"
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
		if pgErr := err.(pg.Error); pgErr != nil && strings.Contains(pgErr.Error(), enums.ErrorDBDuplicatedKeyMsg) {
			return respKit.GenericAlreadyExistsError(enums.ErrorEmailExistsCode, fmt.Sprintf(enums.ErrorEmailExistsMsg, request.Email))
		}

		return respKit.GenericBadRequestError(enums.ErrorInsertCode, err.Error())
	}

	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	u := dto.User{}
	err := r.db.Model(&u).Context(ctx).Where("email = ?", email).Select()
	if err != nil {
		switch err {
		case pg.ErrNoRows:
			return &u, respKit.GenericNotFoundError(enums.ErrorEmailNotFoundCode, fmt.Sprintf(enums.ErrorEmailNotFoundMsg, email))
		default:
			return &u, respKit.GenericBadRequestError(enums.ErrorGetByEmailCode, err.Error())
		}
	}

	return &u, nil
}

func (r *UserRepository) UpdateByEmail(ctx context.Context, request dto.UpdateUserRequest, email string) error {
	err := r.db.Model(&dto.User{}).Context(ctx).Where("email = ?", email).Select()
	if err != nil && err == pg.ErrNoRows {
		return respKit.GenericNotFoundError(enums.ErrorEmailNotFoundCode, fmt.Sprintf(enums.ErrorEmailNotFoundMsg, email))
	}

	if _, err := r.db.Model(&dto.User{}).Context(ctx).Set("name = ?", request.Name).Where("email = ?", email).Update(); err != nil {
		return respKit.GenericBadRequestError(enums.ErrorUpdateCode, err.Error())
	}

	return nil
}

func (r *UserRepository) DeleteByEmail(ctx context.Context, email string) error {
	if _, err := r.db.Model(&dto.User{}).Context(ctx).Where("email = ?", email).Delete(); err != nil {
		switch err {
		case pg.ErrNoRows:
			return respKit.GenericNotFoundError(enums.ErrorEmailNotFoundCode, fmt.Sprintf(enums.ErrorEmailNotFoundMsg, email))
		default:
			return respKit.GenericBadRequestError(enums.ErrorDeleteCode, err.Error())
		}
	}

	return nil
}
