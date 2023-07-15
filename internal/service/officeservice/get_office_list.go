package officeservice

import (
	"context"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) GetOfficeList(ctx context.Context,
	req *customer.GetOfficeListRequest) (*customer.GetOfficeListResponse, error) {

	list, err := s.officeRepository.List(ctx)
	if err != nil {
		s.log.Error("failed get officeList %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	var data []*customer.Office

	for _, office := range list {
		data = append(data, &customer.Office{
			Uuid:      office.Uuid.String(),
			Name:      office.Name,
			Address:   office.Address,
			CreatedAt: timestamppb.New(office.CreatedAt),
		})
	}

	return &customer.GetOfficeListResponse{Result: data}, nil
}
