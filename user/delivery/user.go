package delivery

import (
	"errors"
	"fmt"
	mdl "go-fiber-ticketing/models/user"
	usecase "go-fiber-ticketing/user/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Usecase usecase.Usecase
	logger  *logrus.Logger
}

func Router(router *fiber.App, uc usecase.Usecase, logger *logrus.Logger) {
	u := Handler{Usecase: uc, logger: logger}
	router.Post("api/user", u.GetDataHandler)
	router.Post("api/user/add", u.CreateHandler)
	router.Put("api/user/update", u.UpdateHandler)
	router.Delete("api/user/delete", u.DeleteHandler)
}

func (u Handler) GetDataHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	param := new(mdl.Request)
	var err error

	if err := c.BodyParser(param); err != nil {
		u.logger.Error(err.Error())
		return err
	}
	u.logger.Debug(param)

	result, err := u.Usecase.ReadData(ctx, param)
	if err == nil {
		result.Code = 200
		result.Message = "success retrieve data"
		c.JSON(result)
	} else {
		result.Code = 500
		result.Message = err.Error()
		u.logger.Error(fmt.Sprintf("user.delivery.getDataHandler.Error : %v", err.Error()))
		c.JSON(result)
	}
	return err
}

func (u Handler) CreateHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	param := new(mdl.Request)
	var err error

	if err := c.BodyParser(param); err != nil {
		u.logger.Error(err.Error())
		return err
	}
	u.logger.Debug(param)

	if param.UserId != "" {
		result, err := u.Usecase.CreateData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success create data"
			c.JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			u.logger.Error(fmt.Sprintf("user.delivery.createHandler.Error : %v", err.Error()))
			c.JSON(result)
		}
	} else {
		errn := errors.New("mandatory field is missing")
		c.JSON(mdl.ResponseAll{Code: 400, Message: fmt.Sprintf("user.delivery.createHandler.BadRequest : %v", errn)})
	}
	return err
}

func (u Handler) UpdateHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	param := new(mdl.Request)
	var err error

	if err := c.BodyParser(param); err != nil {
		u.logger.Error(err.Error())
		return err
	}
	u.logger.Debug(param)

	if param.UserId != "" {
		result, err := u.Usecase.UpdateData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success update data"
			c.JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			u.logger.Error(fmt.Sprintf("delivery.user.updateHandler.Error : %v", err.Error()))
			c.JSON(result)
		}
	} else {
		errn := errors.New("mandatory field is missing")
		c.JSON(mdl.ResponseAll{Code: 400, Message: fmt.Sprintf("user.delivery.updateHandler.BadRequest : %v", errn)})
	}
	return err
}

func (u Handler) DeleteHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	param := new(mdl.Request)
	var err error

	if err := c.BodyParser(param); err != nil {
		u.logger.Error(err.Error())
		return err
	}
	u.logger.Debug(param)

	if param.UserId != "" {
		result, err := u.Usecase.DeleteData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success delete data"
			c.JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			u.logger.Error(fmt.Sprintf("user.delivery.deleteHandler.Error : %v", err.Error()))
			c.JSON(result)
		}
	} else {
		errn := errors.New("mandatory field is missing")
		c.JSON(mdl.ResponseAll{Code: 400, Message: fmt.Sprintf("user.delivery.deleteHandler.BadRequest : %v", errn)})
	}
	return err
}
