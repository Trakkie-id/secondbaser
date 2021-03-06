package config

type Config struct {
	DBHost string `json:"DB_HOST"`

	DBUser string `json:"DB_USER"`

	DBPassword string `json:"DB_PASSWORD"`

	DBPort string `json:"DB_PORT"`

	DBDatabase string `json:"DB_DATABASE"`

	GRPCPort string `json:"GRPC_PORT"`

	HTTPPort string `json:"HTTP_PORT"`

	ZipkinEndpoint string `json:"ZIPKIN_ENDPOINT"`

	ApplicationEnv string `json:"APP_ENV"`

	KafkaBrokerAddress string `json:"KAFKA_BROKER_ADDRESS"`

	LogLevel string `json:"LOG_LEVEL"`
}
