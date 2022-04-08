package repository

import (
	"fmt"
	models "go-fiber-ticketing/models/user"

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

const createQuery = "INSERT INTO %s (user_id, username, created_at) VALUES (?,?,NOW())"
const updateQuery = "UPDATE %s SET username = ?, updated_at = NOW() WHERE user_id = ?"
const deleteQuery = "DELETE FROM %s WHERE user_id = ?"
const table = "users"

func (r Repo) GetData(ctx *fasthttp.RequestCtx, param *models.Request) (models.ResponseAll, error) {
	var (
		result []models.Response
		res    models.ResponseAll
		err    error
	)
	query := r.Dbconn.Table(table)
	if param.UserId != "" {
		query = query.Where("user_id = ?", param.UserId)
	}
	err = query.Scan(&result).Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("user.repository.GetData : %v", err.Error()))
		return res, err
	}
	res.Data = result
	return res, err
}

func (r Repo) CreateData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	err = r.Dbconn.Exec(fmt.Sprintf(createQuery, table), param.UserId, param.Username).Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("user.repository.CreateData : %v", err.Error()))
		return err
	}
	return err
}

func (r Repo) UpdateData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(updateQuery, table), param.Username, param.UserId)
	err = query.Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("user.repository.UpdateData : %v", err.Error()))
		return err
	}
	return err
}

func (r Repo) DeleteData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(deleteQuery, table), param.UserId)
	err = query.Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("user.repository.DeleteData : %v", err.Error()))
		return err
	}
	return err
}
