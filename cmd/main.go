package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Aorts/fiber-boiler/config"
	"github.com/Aorts/fiber-boiler/infrastructure"
	"github.com/Aorts/fiber-boiler/internal/adapter/handler/userhdl"
	"github.com/Aorts/fiber-boiler/internal/adapter/repository/userrepo"
	"github.com/Aorts/fiber-boiler/internal/api/middlerware"
	"github.com/Aorts/fiber-boiler/internal/api/routes"
	"github.com/Aorts/fiber-boiler/internal/core/service/authsvc"
	"github.com/Aorts/fiber-boiler/internal/core/service/usersvc"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config.New()
}

func main() {
	mc := infrastructure.NewMongoDB()
	ur := userrepo.NewUserRepository(mc)

	// Service
	usersvc := usersvc.NewUserService(ur)
	authsvc := authsvc.NewAuthService()

	// Handler
	userHandler := userhdl.NewUserHandler(authsvc, usersvc)

	middlewares := middlerware.NewMiddleware()
	app := initFiber()
	// routes
	routesConfig := routes.Config{
		App:         app,
		UserHandler: userHandler,
		Middleware:  middlewares,
	}
	routesConfig.Setup()
	fmt.Println("Server is running on port", config.Get().Server.Port)
	if err := app.Listen(":" + config.Get().Server.Port); err != nil {
		log.Panic(err)
	}
}

func initFiber() *fiber.App {
	middlewares := middlerware.NewMiddleware()
	app := fiber.New(
		fiber.Config{
			ReadTimeout:           5 * time.Second,
			WriteTimeout:          5 * time.Second,
			IdleTimeout:           30 * time.Second,
			DisableStartupMessage: true,
			CaseSensitive:         true,
			StrictRouting:         true,
		},
	)

	app.Use(middlewares.CORSMiddleware())
	return app
}
