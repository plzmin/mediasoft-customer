package officeservice

import (
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"mediasoft-customer/internal/repository/officerepository"
	"mediasoft-customer/pkg/logger"
)

type Service struct {
	log              *logger.Logger
	officeRepository officerepository.OfficeRepository
	customer.UnimplementedOfficeServiceServer
}

func New(log *logger.Logger, officeRepository officerepository.OfficeRepository) *Service {
	return &Service{
		log:              log,
		officeRepository: officeRepository,
	}
}
