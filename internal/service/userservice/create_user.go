package userservice

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mediasoft-customer/internal/model"
	"time"
)

func (s *Service) CreateUser(ctx context.Context,
	req *customer.CreateUserRequest) (*customer.CreateUserResponse, error) {

	if err := req.ValidateAll(); err != nil {
		s.log.Warn("not valid CreateUserRequest %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	uuidOffice, err := uuid.Parse(req.OfficeUuid)
	if err != nil {
		s.log.Error("failed to parse officeUUID %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	user := model.User{
		Uuid:       uuid.New(),
		Name:       req.Name,
		OfficeUuid: uuidOffice,
		CreatedAt:  time.Now(),
	}

	if err = s.userRepository.Create(ctx, &user); err != nil {
		s.log.Error("failed to create user %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &customer.CreateUserResponse{}, nil
}
