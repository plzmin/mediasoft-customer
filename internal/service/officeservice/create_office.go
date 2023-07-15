package officeservice

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mediasoft-customer/internal/model"
)

func (s *Service) CreateOffice(ctx context.Context,
	req *customer.CreateOfficeRequest) (*customer.CreateOfficeResponse, error) {

	if err := req.ValidateAll(); err != nil {
		s.log.Warn("non valid CreateOfficeRequest %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	office := model.Office{
		Uuid:    uuid.New(),
		Name:    req.Name,
		Address: req.Address,
	}

	if err := s.officeRepository.Create(ctx, &office); err != nil {
		s.log.Error("failed create new office %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &customer.CreateOfficeResponse{}, nil
}
