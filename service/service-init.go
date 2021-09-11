package service

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Provide(New())

type IClientService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

type clientService struct {
}

func (c clientService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	panic("implement me")
}

func New() IClientService {
	return &clientService{}
}
