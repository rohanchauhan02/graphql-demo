package user

import (
	"context"

	"github.com/rohanchauhan02/graphql-demo/dto"
	"github.com/rohanchauhan02/graphql-demo/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
}

type Usecase interface {
	Register(ctx context.Context, input dto.CreateUserInput) (*dto.AuthResponse, error)
	Login(ctx context.Context, input dto.LoginInput) (*dto.AuthResponse, error)
	GetUser(ctx context.Context, id uint) (*dto.UserResponse, error)
	UpdateUser(ctx context.Context, id uint, input dto.UpdateUserInput) (*dto.UserResponse, error)
	DeleteUser(ctx context.Context, id uint) error
}
