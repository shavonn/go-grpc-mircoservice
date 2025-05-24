package service

import (
	"context"
	"fmt"
	"time"

	"github.com/shavonn/go-grpc-microservice/pkg/pb"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	// In a real application, you'd inject your database/repository here
	users map[string]*pb.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*pb.User),
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Info().Str("user_id", req.Id).Msg("Getting user")

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "user ID is required")
	}

	user, exists := s.users[req.Id]
	if !exists {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &pb.GetUserResponse{User: user}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Info().Str("email", req.Email).Str("name", req.Name).Msg("Creating user")

	if req.Email == "" || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "email and name are required")
	}

	// Generate a simple ID (in production, use proper UUID)
	userID := fmt.Sprintf("user_%d", time.Now().UnixNano())

	user := &pb.User{
		Id:        userID,
		Email:     req.Email,
		Name:      req.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	s.users[userID] = user

	return &pb.CreateUserResponse{User: user}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	log.Info().Str("user_id", req.Id).Msg("Updating user")

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "user ID is required")
	}

	user, exists := s.users[req.Id]
	if !exists {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	user.UpdatedAt = time.Now().Unix()

	s.users[req.Id] = user

	return &pb.UpdateUserResponse{User: user}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	log.Info().Str("user_id", req.Id).Msg("Deleting user")

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "user ID is required")
	}

	_, exists := s.users[req.Id]
	if !exists {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	delete(s.users, req.Id)

	return &pb.DeleteUserResponse{Success: true}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	log.Info().Int32("page", req.Page).Int32("page_size", req.PageSize).Msg("Listing users")

	// Simple pagination logic
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	var users []*pb.User
	for _, user := range s.users {
		users = append(users, user)
	}

	// Simple pagination (in production, do this at the database level)
	start := int((page - 1) * pageSize)
	end := int(page * pageSize)

	if start >= len(users) {
		users = []*pb.User{}
	} else if end > len(users) {
		users = users[start:]
	} else {
		users = users[start:end]
	}

	return &pb.ListUsersResponse{
		Users:    users,
		Total:    int32(len(s.users)),
		Page:     page,
		PageSize: pageSize,
	}, nil
}
