package configuration

import (
	"time"

	xclients "github.com/syth0le/gopnik/clients"
	xstorage "github.com/syth0le/gopnik/db/postgres"
	xlogger "github.com/syth0le/gopnik/logger"
	xservers "github.com/syth0le/gopnik/servers"
)

const (
	defaultAppName = "dialog-notification"
)

func NewDefaultConfig() *Config {
	return &Config{
		Logger: xlogger.LoggerConfig{
			Level:       xlogger.InfoLevel,
			Encoding:    "console",
			Path:        "stdout",
			Environment: xlogger.Development,
		},
		Application: ApplicationConfig{
			GracefulShutdownTimeout: 15 * time.Second,
			ForceShutdownTimeout:    20 * time.Second,
			App:                     defaultAppName,
		},
		PublicServer: xservers.ServerConfig{
			Enable:   false,
			Endpoint: "",
			Port:     0,
		},
		AdminServer: xservers.ServerConfig{
			Enable:   false,
			Endpoint: "",
			Port:     0,
		},
		InternalGRPCServer: xservers.GRPCServerConfig{
			Port:             0,
			EnableRecover:    false,
			EnableReflection: false,
		},
		Storage: xstorage.StorageConfig{
			EnableMock:            false,
			Hosts:                 []string{},
			Port:                  0,
			Database:              "",
			Username:              "",
			Password:              "",
			SSLMode:               "",
			ConnectionAttempts:    0,
			InitializationTimeout: 5 * time.Second,
		},
		SecondStorage: xstorage.StorageConfig{
			EnableMock:            false,
			Hosts:                 []string{},
			Port:                  0,
			Database:              "",
			Username:              "",
			Password:              "",
			SSLMode:               "",
			ConnectionAttempts:    0,
			InitializationTimeout: 5 * time.Second,
		},
		AuthClient: AuthClientConfig{
			Enable: false,
			Conn: xclients.GRPCClientConnConfig{
				Endpoint:              "",
				UserAgent:             defaultAppName,
				MaxRetries:            0,
				TimeoutBetweenRetries: 0,
				InitTimeout:           0,
				EnableCompressor:      false,
			},
		},
		CounterClient: CounterClientConfig{
			Enable: false,
			Conn: xclients.GRPCClientConnConfig{
				Endpoint:              "",
				UserAgent:             defaultAppName,
				MaxRetries:            0,
				TimeoutBetweenRetries: 0,
				InitTimeout:           0,
				EnableCompressor:      false,
			},
		},
	}
}
