package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	defaultHTTPPort               = "80"
	defaultHTTPReadTimeout        = 15 * time.Second
	defaultHTTPWriteTimeout       = 15 * time.Second
	defaultHTTPIdleTimeout        = 60 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultHTTPSchema             = "https"
)

type (
	Config struct {
		HTTP             HTTPConfig
		POSTGRES         DatabaseConfig
		ExternalServices External
	}

	External struct {
		PaymentServiceURL   string
		ProductServiceURL   string
		WarehouseServiceURL string
	}

	HTTPConfig struct {
		Port               string
		Host               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		IdleTimeout        time.Duration
		MaxHeaderMegabytes int
		Schema             string
	}

	ClientConfig struct {
		Endpoint string
		Username string
		Password string
	}

	DatabaseConfig struct {
		DSN string
	}
)

// New populates Config struct with values from config file
// located at filepath and environment variables.
func New() (cfg Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	godotenv.Load(filepath.Join(root, ".env"))

	httpConfig := HTTPConfig{
		Port:               defaultHTTPPort,
		ReadTimeout:        defaultHTTPReadTimeout,
		WriteTimeout:       defaultHTTPWriteTimeout,
		IdleTimeout:        defaultHTTPIdleTimeout,
		MaxHeaderMegabytes: defaultHTTPMaxHeaderMegabytes,
		Schema:             defaultHTTPSchema,
	}
	cfg.HTTP = httpConfig

	err = envconfig.Process("HTTP", &cfg.HTTP)
	if err != nil {
		return
	}
	err = envconfig.Process("POSTGRES", &cfg.POSTGRES)
	if err != nil {
		return
	}

	cfg.ExternalServices.PaymentServiceURL = os.Getenv("EXTERNAL_PAYMENT_SERVICE_URL")
	cfg.ExternalServices.WarehouseServiceURL = os.Getenv("EXTERNAL_WAREHOUSE_SERVICE_URL")
	cfg.ExternalServices.ProductServiceURL = os.Getenv("EXTERNAL_PRODUCT_SERVICE_URL")

	return
}
