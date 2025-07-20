package grpc

import (
	"context"
	"strconv"

	"github.com/rohanchauhan02/graphql-demo/dto"
	"github.com/rohanchauhan02/graphql-demo/internal/user"
	"github.com/rohanchauhan02/graphql-demo/proto/userpb"
)

type userGRPCServer struct {
	userpb.UnimplementedUserServiceServer
	usecase user.Usecase
}

func NewUserGRPCServer(uc user.Usecase) userpb.UserServiceServer {
	return &userGRPCServer{usecase: uc}
}

func (s *userGRPCServer) Register(ctx context.Context, req *userpb.CreateUserInput) (*userpb.AuthResponse, error) {
	user, err := s.usecase.Register(ctx, dto.CreateUserInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.AuthResponse{
		Token: user.Token,
		User: &userpb.User{
			Id:    strconv.FormatUint(uint64(user.User.ID), 10),
			Name:  user.User.Name,
			Email: user.User.Email,
		},
	}, nil
}

func (s *userGRPCServer) Login(ctx context.Context, req *userpb.LoginInput) (*userpb.AuthResponse, error) {
	user, err := s.usecase.Login(ctx, dto.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.AuthResponse{
		Token: user.Token,
		User: &userpb.User{
			Id:    strconv.FormatUint(uint64(user.User.ID), 10),
			Name:  user.User.Name,
			Email: user.User.Email,
		},
	}, nil
}

func (s *userGRPCServer) GetUser(ctx context.Context, req *userpb.UserIdRequest) (*userpb.User, error) {
	user, err := s.usecase.GetUser(ctx, 1)
	if err != nil {
		return nil, err
	}
	return &userpb.User{
		Id:    strconv.FormatUint(uint64(user.ID), 10),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userGRPCServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.User, error) {
	user, err := s.usecase.UpdateUser(ctx, 1, dto.UpdateUserInput{
		Name:     req.Input.Name,
		Email:    req.Input.Email,
		Password: req.Input.Password,
	})
	if err != nil {
		return nil, err
	}
	return &userpb.User{
		Id:    strconv.FormatUint(uint64(user.ID), 10),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userGRPCServer) DeleteUser(ctx context.Context, req *userpb.UserIdRequest) (*userpb.DeleteUserResponse, error) {
	err := s.usecase.DeleteUser(ctx, 1)
	return &userpb.DeleteUserResponse{Success: err == nil}, err
}
