package repository

import (
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewRepository))

type IRepository interface {
	AddClient()
}

type repository struct {
}

func (r repository) AddClient() {

}

func NewRepository() IRepository {
	return &repository{}
}
