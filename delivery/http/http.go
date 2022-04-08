package http

import (
	checkout_router "go-fiber-ticketing/checkout/delivery"
	checkout "go-fiber-ticketing/checkout/usecase"
	ticket_router "go-fiber-ticketing/ticket/delivery"
	ticket "go-fiber-ticketing/ticket/usecase"
	user_router "go-fiber-ticketing/user/delivery"
	user "go-fiber-ticketing/user/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

type HttpRouter struct {
	UserUsecase     user.Usecase
	TicketUsecase   ticket.Usecase
	CheckoutUsecase checkout.Usecase
}

func InitRouter(usecase *HttpRouter, logger *logrus.Logger) *fiber.App {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Headers, Accept-Encoding, X-CSRF-Token",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	user_router.Router(router, usecase.UserUsecase, logger)
	checkout_router.Router(router, usecase.CheckoutUsecase, logger)
	ticket_router.Router(router, usecase.TicketUsecase, logger)

	return router
}
