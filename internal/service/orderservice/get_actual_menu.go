package orderservice

import (
	"context"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetActualMenu(ctx context.Context,
	req *customer.GetActualMenuRequest) (*customer.GetActualMenuResponse, error) {
	res, err := s.restaurant.GetActualMenu(ctx)
	if err != nil {
		s.log.Error("failed to GetActualMenu %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &customer.GetActualMenuResponse{
		Salads:    restaurantToCustomerProduct(res.Menu.Salads),
		Garnishes: restaurantToCustomerProduct(res.Menu.Garnishes),
		Meats:     restaurantToCustomerProduct(res.Menu.Meats),
		Soups:     restaurantToCustomerProduct(res.Menu.Soups),
		Drinks:    restaurantToCustomerProduct(res.Menu.Drinks),
		Desserts:  restaurantToCustomerProduct(res.Menu.Desserts),
	}, nil
}

func restaurantToCustomerProduct(pl []*restaurant.Product) []*customer.Product {
	var cpl []*customer.Product
	for _, p := range pl {
		cpl = append(cpl, &customer.Product{
			Uuid:        p.Uuid,
			Name:        p.Name,
			Description: p.Description,
			Type:        customer.CustomerProductType(p.Type),
			Weight:      p.Weight,
			Price:       p.Price,
			CreatedAt:   p.CreatedAt,
		})
	}
	return cpl
}
