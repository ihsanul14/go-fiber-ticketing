package repository

import (
	"fmt"
	models "go-fiber-ticketing/models/checkout"

	"github.com/rs/xid"
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
	CreateData(ctx *fasthttp.RequestCtx, request *models.CreateRequest) error
	DeleteData(ctx *fasthttp.RequestCtx, request *models.Request) error
	GetSummaryData(ctx *fasthttp.RequestCtx, request *models.Request) (res models.ResponseSummaryAll, err error)
	Payment(ctx *fasthttp.RequestCtx, request *models.Request) error
}

func NewRepository(dbconn *gorm.DB, logger *logrus.Logger) Repository {
	return &Repo{Dbconn: dbconn, logger: logger}
}

const createQuery = "INSERT INTO %s (checkout_id, user_id, ticket_id, is_purchased, created_at) VALUES (?,?,?,0,NOW())"
const deleteQuery = "DELETE FROM %s WHERE id = ?"
const paymentQuery = "UPDATE %s SET is_purchased = 1, updated_at = NOW() WHERE checkout_id = ?"

const summaryQuery = `
	SELECT a.user_id user_id, a.username username, b.checkout_id checkout_id, c.ticket_id ticket_id, c.acara acara, SUM(c.harga) harga
	FROM users a
	INNER JOIN checkout b
	ON a.user_id = b.user_id
	INNER JOIN ticket c
	ON b.ticket_id = c.ticket_id
	WHERE a.user_id = ? and b.checkout_id = ? and b.is_purchased = ?
	GROUP BY a.user_id, b.checkout_id, c.ticket_id
`

const table = "checkout"

func (r Repo) GetData(ctx *fasthttp.RequestCtx, param *models.Request) (models.ResponseAll, error) {
	var (
		result []models.Response
		res    models.ResponseAll
		err    error
	)
	query := r.Dbconn.Table(table)
	if param.Id != 0 {
		query = query.Where("id = ?", param.Id)
	}
	if param.UserId != "" {
		query = query.Where("user_id = ?", param.UserId)
	}
	if param.TicketId != "" {
		query = query.Where("ticket_id = ?", param.TicketId)
	}
	if param.CheckoutId != "" {
		query = query.Where("checkout_id = ?", param.CheckoutId)
	}
	err = query.Scan(&result).Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("repository.repository.GetData : %v", err.Error()))
		return res, err
	}
	res.Data = result
	return res, err
}

func (r Repo) CreateData(ctx *fasthttp.RequestCtx, param *models.CreateRequest) (err error) {
	if param.CheckoutId == "" {
		param.CheckoutId = xid.New().String()
	}
	for _, val := range param.TicketId {
		err = r.Dbconn.Exec(fmt.Sprintf(createQuery, table), param.CheckoutId, param.UserId, val).Error
		if err != nil {
			r.logger.Error(fmt.Sprintf("repository.repository.CreateData : %v", err.Error()))
			return err
		}
	}
	return err
}

func (r Repo) DeleteData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(deleteQuery, table), param.Id)
	err = query.Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("repository.repository.DeleteData : %v", err.Error()))
		return err
	}
	return err
}

func (r Repo) GetSummaryData(ctx *fasthttp.RequestCtx, param *models.Request) (models.ResponseSummaryAll, error) {
	var (
		result []models.ResponseSummary
		res    models.ResponseSummaryAll
		err    error
	)
	err = r.Dbconn.Raw(summaryQuery, param.UserId, param.CheckoutId, param.IsPurchased).Scan(&result).Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("repository.repository.GetSummaryData : %v", err.Error()))
		return res, err
	}
	res.Data = result
	return res, err
}

func (r Repo) Payment(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(paymentQuery, table), param.CheckoutId)
	err = query.Error
	if err != nil {
		r.logger.Error(fmt.Sprintf("repository.repository.UpdateData : %v", err.Error()))
		return err
	}
	return err
}
