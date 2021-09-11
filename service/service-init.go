package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

type IClientService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

type clientService struct {

}

func (c clientService) Add(ctx context.Context, numA, numB float32) (float32, error) {
	panic("implement me")
}

func New(looger log.Logger) IClientService{
	return &clientService{

	}
}