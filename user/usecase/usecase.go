package usecase

import (
	"fmt"
	mdl "go-fiber-ticketing/models/user"
	repo "go-fiber-ticketing/user/repository"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type UsecaseModul struct {
	Repo   repo.Repository
	logger *logrus.Logger
}

type Usecase interface {
	ReadData(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseAll, err error)
	CreateData(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseAll, err error)
	UpdateData(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseAll, err error)
	DeleteData(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseAll, err error)
}

func NewUsecase(u repo.Repository, logger *logrus.Logger) Usecase {
	return &UsecaseModul{Repo: u, logger: logger}
}
func (u UsecaseModul) ReadData(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseAll, error) {
	res, err := u.Repo.GetData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("user.usecase.ReadData: %v", err.Error()))
		return res, err
	}
	return res, err
}

func (u *UsecaseModul) CreateData(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseAll, error) {
	var res mdl.ResponseAll
	err := u.Repo.CreateData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("user.usecase.CreateData: %v", err.Error()))
		return res, err
	}
	return res, err
}

func (u *UsecaseModul) UpdateData(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseAll, error) {
	var res mdl.ResponseAll
	err := u.Repo.UpdateData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("user.usecase.UpdateData: %v", err.Error()))
		return res, err
	}
	return res, err
}

func (u *UsecaseModul) DeleteData(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseAll, error) {
	var res mdl.ResponseAll
	err := u.Repo.DeleteData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("user.usecase.DeleteData: %v", err.Error()))
		return res, err
	}
	return res, err
}
