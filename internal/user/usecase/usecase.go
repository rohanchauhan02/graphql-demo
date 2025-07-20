package usecase

import (
	"context"
	"errors"

	"github.com/rohanchauhan02/graphql-demo/dto"
	"github.com/rohanchauhan02/graphql-demo/internal/user"
	"github.com/rohanchauhan02/graphql-demo/models"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return &userUsecase{userRepo: userRepo}
}

func (uc *userUsecase) Register(ctx context.Context, input dto.CreateUserInput) (*dto.AuthResponse, error) {
	_, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token := "generated-jwt-token"

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}, nil
}

func (uc *userUsecase) Login(ctx context.Context, input dto.LoginInput) (*dto.AuthResponse, error) {
	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token := "generated-jwt-token"

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}, nil
}

func (uc *userUsecase) GetUser(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (uc *userUsecase) UpdateUser(ctx context.Context, id uint, input dto.UpdateUserInput) (*dto.UserResponse, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (uc *userUsecase) DeleteUser(ctx context.Context, id uint) error {
	return uc.userRepo.Delete(ctx, id)
}
