package config

type Config struct {
	HTTP struct {
		IP   string `yaml:"IP" env:"HTTP_IP" env-default:"localhost"`
		Port int    `yaml:"Port" env:"HTTP_PORT" env-default:"8000"`
	} `yaml:"HTTP"`
	GRPC struct {
		IP   string `yaml:"IP" env:"GRPC_IP" env-default:"localhost"`
		Port int    `yaml:"Port" env:"GRPC_PORT" env-default:"8001"`
	} `yaml:"GRPC"`

	PostgresSQL struct {
		Username string `yaml:"Username" env:"PG_USER"  env-default:"postgres"`
		Password string `yaml:"Password" env:"PG_PWD" env-default:"postgres"`
		Host     string `yaml:"Host" env:"PG_HOST"  env-default:"localhost"`
		Port     string `yaml:"Port" env:"PG_PORT" env-default:"5432"`
		Database string `yaml:"Database" env:"PG_DATABASE"  env-default:"customer_db"`
	} `yaml:"PostgresSQL"`

	RestaurantGRPC struct {
		IP   string `yaml:"IP" env:"RESTAURANT_IP" env-default:"localhost"`
		Port int    `yaml:"Port" env:"RESTAURANT_PORT" env-default:"8003"`
	} `yaml:"RestaurantGRPC"`

	Kafka []string `yaml:"Kafka" env:"KAFKA" env-default:"localhost:9092"`
}
