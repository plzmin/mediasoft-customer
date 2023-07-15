package orderservice

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mediasoft-customer/internal/model"
	"time"
)

func (s *Service) CreateOrder(ctx context.Context, req *customer.CreateOrderRequest) (*customer.CreateOrderResponse, error) {
	if err := req.ValidateAll(); err != nil {
		s.log.Warn("not valid CreateOrderRequest %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, err := s.restaurant.GetActualMenu(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	openTime := res.Menu.OpeningRecordAt.AsTime()
	closeTime := res.Menu.ClosingRecordAt.AsTime()
	if openTime.Before(time.Now()) && closeTime.After(time.Now()) {
		return nil, status.Error(codes.Canceled, "closed record")
	}

	userUuid, err := uuid.Parse(req.UserUuid)
	if err != nil {
		s.log.Error("failed to parse userUUID %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	order := model.Order{
		Uuid:      uuid.New(),
		UserUuid:  userUuid,
		Salads:    req.Salads,
		Garnishes: req.Garnishes,
		Meats:     req.Meats,
		Soups:     req.Soups,
		Drinks:    req.Drinks,
		Desserts:  req.Desserts,
	}

	if err = s.orderRepository.Create(ctx, &order); err != nil {
		s.log.Error("failed to create order %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = s.producer.SendMessage("order", order); err != nil {
		s.log.Error("failed send message to broker %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &customer.CreateOrderResponse{}, nil
}
