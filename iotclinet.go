package main

import (
	"go.uber.org/fx"

	"iotclient/service"
)

func main() {
	fx.New(fx.Provide(service.New())).
		Run()
}
