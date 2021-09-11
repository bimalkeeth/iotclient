package service

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewClient))

type IClientService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

type clientService struct {
}

func (c clientService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return 10, nil
}

func NewClient() IClientService {
	return clientService{}
}
