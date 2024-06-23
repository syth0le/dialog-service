package application

import (
	xservers "github.com/syth0le/gopnik/servers"

	"github.com/syth0le/dialog-service/internal/handler/internalapi"
	inpb "github.com/syth0le/dialog-service/proto/internalapi"
)

func (a *App) newInternalGRPCServer(env *env) *xservers.GRPCServer {
	server := xservers.NewGRPCServer(
		a.Config.InternalGRPCServer,
		a.Logger,
		xservers.GRPCWithServerName("internal grpc api"),
	)

	inpb.RegisterDialogServiceServer(server.Server, &internalapi.DialogHandler{DialogService: env.dialogService})

	return server
}
