package main

import (
	cr "go-fiber-ticketing/checkout/repository"
	cu "go-fiber-ticketing/checkout/usecase"
	"go-fiber-ticketing/database"
	delivery_http "go-fiber-ticketing/delivery/http"
	tr "go-fiber-ticketing/ticket/repository"
	tu "go-fiber-ticketing/ticket/usecase"
	ur "go-fiber-ticketing/user/repository"
	uu "go-fiber-ticketing/user/usecase"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	var baseLogger = logrus.New()

	baseLogger.Formatter = &logrus.JSONFormatter{}

	if os.Getenv("environment") == "development" {
		baseLogger.SetLevel(logrus.DebugLevel)
	} else {
		baseLogger.SetLevel(logrus.InfoLevel)
	}

	mysqlConn, err := database.ConnectMySQL(baseLogger)
	if err != nil {
		return
	}
	userRepository := ur.NewRepository(mysqlConn, baseLogger)
	userUsecase := uu.NewUsecase(userRepository, baseLogger)

	checkoutRepository := cr.NewRepository(mysqlConn, baseLogger)
	checkoutUsecase := cu.NewUsecase(checkoutRepository, baseLogger)

	ticketRepository := tr.NewRepository(mysqlConn, baseLogger)
	ticketUsecase := tu.NewUsecase(ticketRepository, baseLogger)

	httpRouter := &delivery_http.HttpRouter{UserUsecase: userUsecase, CheckoutUsecase: checkoutUsecase, TicketUsecase: ticketUsecase}
	router := delivery_http.InitRouter(httpRouter, baseLogger)
	router.Use(logger.New())

	router.Listen(":" + os.Getenv("PORT"))
}
