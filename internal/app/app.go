package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mediasoft-customer/internal/bootstrap"
	"mediasoft-customer/internal/client"
	"mediasoft-customer/internal/config"
	"mediasoft-customer/internal/kafka"
	"mediasoft-customer/internal/repository/officerepository/officesqlx"
	"mediasoft-customer/internal/repository/orderrepository/ordersqlx"
	"mediasoft-customer/internal/repository/userrepository/usersqlx"
	"mediasoft-customer/internal/service/officeservice"
	"mediasoft-customer/internal/service/orderservice"
	"mediasoft-customer/internal/service/userservice"
	"mediasoft-customer/pkg/logger"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(log *logger.Logger, cfg config.Config) error {
	s := grpc.NewServer()
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	go runGRPCServer(log, cfg, s)
	go runHTTPServer(log, ctx, cfg, mux)

	gracefulShutDown(log, s, cancel)

	return nil
}

func runGRPCServer(log *logger.Logger, cfg config.Config, s *grpc.Server) {
	db, err := bootstrap.InitSqlxDB(cfg)
	if err != nil {
		log.Fatal("failed init db conn %v", err)
	}

	restaurant, err := client.New(cfg)
	if err != nil {
		log.Fatal("failed init restaurant client %v", err)
	}
	defer func(restaurant *client.RestaurantClient) {
		err = restaurant.Close()
		if err != nil {
			log.Fatal("failed close restaurant conn %v", err)
		}
	}(restaurant)

	producer, err := kafka.New(cfg.Kafka)
	if err != nil {
		log.Fatal("failed init kafka conn %v", err)
	}
	defer func(producer *kafka.Producer) {
		err = producer.Close()
		if err != nil {
			log.Fatal("failed close kafka conn %v", err)
		}
	}(producer)

	officeService := officeservice.New(log, officesqlx.New(db))
	orderService := orderservice.New(log, ordersqlx.New(db), producer, restaurant)
	userService := userservice.New(log, usersqlx.New(db))
	customer.RegisterOfficeServiceServer(s, officeService)
	customer.RegisterOrderServiceServer(s, orderService)
	customer.RegisterUserServiceServer(s, userService)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.IP, cfg.GRPC.Port))
	if err != nil {
		log.Fatal("failed to listen tcp %s:%d, %v", cfg.GRPC.IP, cfg.GRPC.Port, err)
	}

	log.Info("starting listening grpc server at %s:%d", cfg.GRPC.IP, cfg.GRPC.Port)
	if err = s.Serve(l); err != nil {
		log.Fatal("error service grpc server %v", err)
	}

}

func runHTTPServer(log *logger.Logger, ctx context.Context, cfg config.Config, mux *runtime.ServeMux) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endPoint := fmt.Sprintf("%s:%d", cfg.HTTP.IP, cfg.GRPC.Port)

	if err := customer.RegisterOfficeServiceHandlerFromEndpoint(ctx, mux, endPoint, dialOptions); err != nil {
		log.Fatal("failed to register office service %v", err)
	}

	if err := customer.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, endPoint, dialOptions); err != nil {
		log.Fatal("failed to register order service %v", err)
	}

	if err := customer.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endPoint, dialOptions); err != nil {
		log.Fatal("failed to register user service %v", err)
	}

	log.Info("starting listening http server at %s:%d", cfg.HTTP.IP, cfg.HTTP.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.HTTP.IP, cfg.HTTP.Port), mux); err != nil {
		log.Fatal("error service http server %v", err)
	}

}

func gracefulShutDown(log *logger.Logger, s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)
	sig := <-ch
	log.Info("Received shutdown signal: %v -  Graceful shutdown done", sig)
	s.GracefulStop()
	cancel()
}
