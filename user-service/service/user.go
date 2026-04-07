package service

import (
	"context"
	"database/sql"
	"user-service/auth"
	pb "user-service/proto"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	DB *sql.DB
}

// Create User
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {

	hashed, _ := auth.HashPassword(req.Password)

	var id int32
	err := s.DB.QueryRow(
		"INSERT INTO grpcUsers(name,email,password) VALUES ($1,$2,$3) RETURNING id",
		req.Name, req.Email, hashed,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

// Login
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	var id int32
	var password string

	err := s.DB.QueryRow(
		"SELECT id,password FROM users WHERE email=$1",
		req.Email,
	).Scan(&id, &password)
	if err != nil {
		return nil, err
	}

	if !auth.CheckPassword(password, req.Password) {
		return nil, err
	}

	token, _ := auth.GenerateJWT(int(id))

	return &pb.LoginResponse{
		Token: token,
	}, nil
}
