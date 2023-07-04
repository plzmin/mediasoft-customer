package userservice

import (
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"mediasoft-customer/internal/repository/userrepository"
	"mediasoft-customer/pkg/logger"
)

type Service struct {
	log            *logger.Logger
	userRepository userrepository.UserRepository
	customer.UnimplementedUserServiceServer
}

func New(log *logger.Logger, userRepository userrepository.UserRepository) *Service {
	return &Service{
		log:            log,
		userRepository: userRepository,
	}
}
