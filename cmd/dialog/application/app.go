package application

import (
	"context"
	"fmt"
	"syscall"

	xclients "github.com/syth0le/gopnik/clients"
	xcloser "github.com/syth0le/gopnik/closer"
	"go.uber.org/zap"

	"github.com/syth0le/dialog-service/cmd/dialog/configuration"
	"github.com/syth0le/dialog-service/internal/clients/auth"
	"github.com/syth0le/dialog-service/internal/service/dialog"
	"github.com/syth0le/dialog-service/internal/storage/postgres"
)

type App struct {
	Config *configuration.Config
	Logger *zap.Logger
	Closer *xcloser.Closer
}

func New(cfg *configuration.Config, logger *zap.Logger) *App {
	return &App{
		Config: cfg,
		Logger: logger,
		Closer: xcloser.NewCloser(logger, cfg.Application.GracefulShutdownTimeout, cfg.Application.ForceShutdownTimeout, syscall.SIGINT, syscall.SIGTERM),
	}
}

func (a *App) Run() error {
	ctx, cancelFunction := context.WithCancel(context.Background())
	a.Closer.Add(func() error {
		cancelFunction()
		return nil
	})

	envStruct, err := a.constructEnv(ctx)
	if err != nil {
		return fmt.Errorf("construct env: %w", err)
	}

	internalGRPCServer := a.newInternalGRPCServer(envStruct)
	a.Closer.AddForce(internalGRPCServer.ForcefullyStop)
	a.Closer.Add(internalGRPCServer.GracefullyStop)

	a.Closer.Run(internalGRPCServer.Run)

	httpServer := a.newHTTPServer(envStruct)
	a.Closer.Add(httpServer.GracefulStop()...)

	a.Closer.Run(httpServer.Run()...)
	a.Closer.Wait()
	return nil
}

type env struct {
	authClient    auth.Client
	dialogService dialog.Service
}

func (a *App) constructEnv(ctx context.Context) (*env, error) {
	postgresDB, err := postgres.NewStorage(a.Logger, a.Config.Storage)
	if err != nil {
		return nil, fmt.Errorf("new storage: %w", err)
	}
	a.Closer.Add(postgresDB.Close)

	authClient, err := a.makeAuthClient(ctx, a.Config.AuthClient)
	if err != nil {
		return nil, fmt.Errorf("make auth client: %w", err)
	}

	return &env{
		authClient:    authClient,
		dialogService: dialog.NewServiceImpl(a.Logger, postgresDB),
	}, nil
}

func (a *App) makeAuthClient(ctx context.Context, cfg configuration.AuthClientConfig) (auth.Client, error) {
	if !cfg.Enable {
		return auth.NewClientMock(a.Logger), nil
	}

	connection, err := xclients.NewGRPCClientConn(ctx, cfg.Conn)
	if err != nil {
		return nil, fmt.Errorf("new grpc conn: %w", err)
	}

	a.Closer.Add(connection.Close)

	return auth.NewAuthImpl(a.Logger, connection), nil
}
