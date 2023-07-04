package client

import (
	"context"
	"fmt"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mediasoft-customer/internal/config"
	"time"
)

type RestaurantClient struct {
	client restaurant.MenuServiceClient
}

func New(cfg config.Config) (*RestaurantClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.RestaurantGRPC.IP, cfg.RestaurantGRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := restaurant.NewMenuServiceClient(conn)

	return &RestaurantClient{client: client}, err
}

func (r *RestaurantClient) GetActualMenu(ctx context.Context) (*restaurant.GetMenuResponse, error) {
	res, err := r.client.GetMenu(ctx, &restaurant.GetMenuRequest{
		OnDate: timestamppb.New(time.Now()),
	})
	if err != nil {
		return nil, err
	}
	return res, err
}

func (r *RestaurantClient) Close() error {
	return r.Close()
}
