package orderservice

import (
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"mediasoft-customer/internal/client"
	"mediasoft-customer/internal/kafka"
	"mediasoft-customer/internal/repository/orderrepository"
	"mediasoft-customer/pkg/logger"
)

type Service struct {
	log             *logger.Logger
	orderRepository orderrepository.OrderRepository
	producer        *kafka.Producer
	restaurant      *client.RestaurantClient
	customer.UnimplementedOrderServiceServer
}

func New(log *logger.Logger,
	orderRepository orderrepository.OrderRepository,
	producer *kafka.Producer,
	restaurant *client.RestaurantClient) *Service {
	return &Service{
		log:             log,
		orderRepository: orderRepository,
		producer:        producer,
		restaurant:      restaurant,
	}
}
