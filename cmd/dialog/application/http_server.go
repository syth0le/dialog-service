package application

import (
	"github.com/go-chi/chi/v5"
	xservers "github.com/syth0le/gopnik/servers"

	"github.com/syth0le/dialog-service/internal/handler/publicapi"
)

func (a *App) newHTTPServer(env *env) *xservers.HTTPServerWrapper {
	return xservers.NewHTTPServerWrapper(
		a.Logger,
		xservers.WithAdminServer(a.Config.AdminServer),
		xservers.WithPublicServer(a.Config.PublicServer, a.publicMux(env)),
	)
}

func (a *App) publicMux(env *env) *chi.Mux {
	mux := chi.NewMux()

	handler := &publicapi.Handler{
		Logger:        a.Logger,
		DialogService: env.dialogService,
	}

	mux.Route("/dialog", func(r chi.Router) {
		r.Use(env.authClient.AuthenticationInterceptor)

		r.Post("/", handler.CreateDialog) // todo: make group dialogs
		r.Post("/send", handler.CreateMessage)
		r.Get("/{dialogID}/list", handler.GetDialogMessages)
	})

	return mux
}
