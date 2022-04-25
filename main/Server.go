package main

import (
	"github.com/gin-gonic/gin"
	"payment/Delivery/api"
	"payment/Delivery/middleware"
	"payment/config"
	"payment/util"
)

type AppServer interface {
	Run()
}

type appServer struct {
	r *gin.Engine
	c config.Config
}

func (a *appServer) initHandler() {
	a.r.Use(middleware.AuthTokenMiddleware())
	a.v1()
}

func (a *appServer) v1() {
	customerGroup := a.r.Group("/customers")
	paymentGroup := a.r.Group("/payment")
	api.LoginApiRoute(customerGroup, a.c.UseCaseManager.CustomerUseCase())
	api.PaymentApi(paymentGroup, a.c.UseCaseManager.TransferUseCase())
}

func (a *appServer) Run() {
	a.initHandler()
	err := a.r.Run(a.c.ApiConfig.Url)
	if err != nil {
		util.Log.Fatal().Msg("Server Failed to run")
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig(".", "config")

	return &appServer{
		r: r,
		c: c,
	}
}
