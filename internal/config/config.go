package config

type Config struct {
	HTTP struct {
		IP   string `env:"HTTP_IP" env-default:"localhost"`
		Port int    `env:"HTTP_PORT" env-default:"8000"`
	}
	GRPC struct {
		IP   string `env:"GRPC_IP" env-default:"localhost"`
		Port int    `env:"GRPC_PORT" env-default:"8001"`
	}

	PostgresSQL struct {
		Username string `env:"PG_USER"  env-default:"postgres"`
		Password string `env:"PG_PWD" env-default:"postgres"`
		Host     string `env:"PG_HOST"  env-default:"localhost"`
		Port     string `env:"PG_PORT" env-default:"5432"`
		Database string `env:"PG_DATABASE"  env-default:"customer_db"`
	}

	RestaurantGRPC struct {
		IP   string `env:"RESTAURANT_IP" env-default:"localhost"`
		Port int    `env:"RESTAURANT_PORT" env-default:"8003"`
	}

	Kafka []string `env:"KAFKA" env-default:"localhost:9092"`
}
