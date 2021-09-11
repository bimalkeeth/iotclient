package service

import (
	"context"
	"fmt"

	"go.uber.org/fx"

	"iotclient/repository"
)

var Module = fx.Options(fx.Provide(NewClient))

type IClientService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

type clientService struct {
	repository repository.IRepository
}

func (c clientService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	ss := c
	fmt.Print(ss)
	return 10, nil
}

func NewClient(repository repository.IRepository) IClientService {
	return clientService{
		repository,
	}
}
