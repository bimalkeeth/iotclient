package main

import (
	"context"
	"net/http"

	"go.uber.org/fx"

	"iotclient/service"
)

func main() {
	app := fx.New(service.Module, fx.Provide(http.NewServeMux), fx.Invoke(registerHooks))
	app.Run()
}

func registerHooks(
	lifecycle fx.Lifecycle, mux *http.ServeMux,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				go func() {
					err := http.ListenAndServe(":8080", mux)
					if err != nil {

					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}
