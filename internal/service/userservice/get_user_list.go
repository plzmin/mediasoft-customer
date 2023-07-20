package userservice

import (
	"context"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mediasoft-customer/internal/model"
)

func (s *Service) GetUserList(ctx context.Context,
	req *customer.GetUserListRequest) (*customer.GetUserListResponse, error) {

	var list []*model.User
	list, err := s.userRepository.List(ctx, req.OfficeUuid)
	if err != nil {
		s.log.Error("failed get userList %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	var data []*customer.User
	for _, u := range list {
		data = append(data, &customer.User{
			Uuid:       u.Uuid.String(),
			Name:       u.Name,
			OfficeUuid: u.OfficeUuid.String(),
			OfficeName: u.OfficeName,
			CreatedAt:  timestamppb.New(u.CreatedAt),
		})

	}

	return &customer.GetUserListResponse{Result: data}, nil
}
