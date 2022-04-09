package usecase

import (
	"errors"
	"fmt"
	repo "go-fiber-ticketing/checkout/repository"
	mdl "go-fiber-ticketing/models/checkout"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type UsecaseModul struct {
	Repo   repo.Repository
	logger *logrus.Logger
}

type Usecase interface {
	ReadData(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseAll, err error)
	CreateData(ctx *fasthttp.RequestCtx, request *mdl.CreateRequest) (res mdl.ResponseAll, err error)
	DeleteData(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseAll, err error)
	ReadSummaryData(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseSummaryAll, err error)
	Payment(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseAll, err error)
}

func NewUsecase(u repo.Repository, logger *logrus.Logger) Usecase {
	return &UsecaseModul{Repo: u, logger: logger}
}

func (u UsecaseModul) ReadData(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseAll, error) {
	res, err := u.Repo.GetData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("usecase.usecase.ReadData: %v", err.Error()))
		return res, err
	}
	return res, err
}

func (u *UsecaseModul) CreateData(ctx *fasthttp.RequestCtx, param *mdl.CreateRequest) (mdl.ResponseAll, error) {
	var res mdl.ResponseAll
	err := u.Repo.CreateData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("usecase.usecase.CreateData: %v", err.Error()))
		return res, err
	}
	return res, err
}

func (u *UsecaseModul) DeleteData(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseAll, error) {
	var res mdl.ResponseAll
	err := u.Repo.DeleteData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("usecase.usecase.DeleteData: %v", err.Error()))
		return res, err
	}
	return res, err
}

func (u UsecaseModul) ReadSummaryData(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseSummaryAll, error) {
	res, err := u.Repo.GetSummaryData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("usecase.usecase.ReadData: %v", err.Error()))
		return res, err
	}

	for _, val := range res.Data {
		res.TotalHarga = res.TotalHarga + val.Harga
	}

	return res, err
}

func (u *UsecaseModul) Payment(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseAll, error) {
	var res mdl.ResponseAll
	p, err := u.ReadSummaryData(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("usecase.usecase.Payment: %v", err.Error()))
		return res, err
	}
	u.logger.Info("checking user_id")
	if param.UserId != p.Data[0].UserId {
		errv := errors.New("user_id invalid")
		u.logger.Error(fmt.Sprintf("usecase.usecase.Payment: %v", errv.Error()))
		return res, errv
	}
	u.logger.Info("checking payment account")
	if param.PaymentAccount != p.TotalHarga {
		errv := errors.New("payment account invalid")
		u.logger.Error(fmt.Sprintf("usecase.usecase.Payment: %v", errv.Error()))
		return res, errv
	}
	err = u.Repo.Payment(ctx, param)
	if err != nil {
		u.logger.Error(fmt.Sprintf("usecase.usecase.Payment: %v", err.Error()))
		return res, err
	}
	u.logger.Info("success payment")
	return res, err
}
