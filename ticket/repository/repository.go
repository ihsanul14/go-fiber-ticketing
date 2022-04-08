package repository

import (
	"fmt"
	models "go-fiber-ticketing/models/ticket"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

type Repo struct {
	Dbconn *gorm.DB
	logger *logrus.Logger
}

type Repository interface {
	GetData(ctx *fasthttp.RequestCtx, request *models.Request) (res models.ResponseAll, err error)
	CreateData(ctx *fasthttp.RequestCtx, request *models.Request) error
	UpdateData(ctx *fasthttp.RequestCtx, request *models.Request) error
	DeleteData(ctx *fasthttp.RequestCtx, request *models.Request) error
}

func NewRepository(dbconn *gorm.DB, logger *logrus.Logger) Repository {
	return &Repo{Dbconn: dbconn, logger: logger}
}

const createQuery = "INSERT INTO %s (ticket_id, acara, harga, created_at) VALUES (?,?,?,NOW())"
const updateQuery = "UPDATE %s SET acara = ?, harga = ?, updated_at = NOW() WHERE ticket_id = ?"
const deleteQuery = "DELETE FROM %s WHERE ticket_id = ?"
const table = "ticket"

func (r Repo) GetData(ctx *fasthttp.RequestCtx, param *models.Request) (models.ResponseAll, error) {
	var (
		result []models.Response
		res    models.ResponseAll
		err    error
	)
	query := r.Dbconn.Table(table)
	if param.TicketId != "" {
		query = query.Where("ticket_id = ?", param.TicketId)
	}
	err = query.Scan(&result).Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("repository.repository.GetData : %v", err.Error()))
		return res, err
	}
	res.Data = result
	return res, err
}

func (r Repo) CreateData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	err = r.Dbconn.Exec(fmt.Sprintf(createQuery, table), param.TicketId, param.Acara, param.Harga).Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("repository.repository.CreateData : %v", err.Error()))
		return err
	}
	return err
}

func (r Repo) UpdateData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(updateQuery, table), param.Acara, param.Harga, param.TicketId)
	err = query.Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("repository.repository.UpdateData : %v", err.Error()))
		return err
	}
	return err
}

func (r Repo) DeleteData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(deleteQuery, table), param.TicketId)
	err = query.Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("repository.repository.DeleteData : %v", err.Error()))
		return err
	}
	return err
}
