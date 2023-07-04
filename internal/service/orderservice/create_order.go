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

	salads, err := customerOrderItemToModel(req.Salads)
	if err != nil {
		s.log.Error("failed to parse salads uuid %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	garnishes, err := customerOrderItemToModel(req.Garnishes)
	if err != nil {
		s.log.Error("failed to parse garnishes uuid %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	meats, err := customerOrderItemToModel(req.Meats)
	if err != nil {
		s.log.Error("failed to parse meats uuid %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	soups, err := customerOrderItemToModel(req.Soups)
	if err != nil {
		s.log.Error("failed to parse soups uuid %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	drinks, err := customerOrderItemToModel(req.Drinks)
	if err != nil {
		s.log.Error("failed to parse drinks uuid %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	desserts, err := customerOrderItemToModel(req.Desserts)
	if err != nil {
		s.log.Error("failed to parse desserts uuid %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	orderUuid := uuid.New()
	order := model.Order{
		Uuid:      orderUuid,
		UserUuid:  userUuid,
		Salads:    salads,
		Garnishes: garnishes,
		Meats:     meats,
		Soups:     soups,
		Drinks:    drinks,
		Desserts:  desserts,
		CreatedAt: time.Now(),
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

func customerOrderItemToModel(orderItems []*customer.OrderItem) ([]*model.OrderItem, error) {
	var orl []*model.OrderItem
	for _, orderItem := range orderItems {
		productUuid, err := uuid.Parse(orderItem.ProductUuid)
		if err != nil {
			return nil, err
		}
		orl = append(orl, &model.OrderItem{
			Count:       orderItem.Count,
			ProductUuid: productUuid,
		})
	}
	return orl, nil
}
