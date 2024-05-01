package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/config"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/delivery"
	"github.com/thoriqulumar/cats-social-service-w1/internal/middleware"
)

func initRouter(h *delivery.Handler) {
	r := gin.Default()

	// registerRouters(app)
	registerRouters(r, h)

	// TODO: graceful shutdown
	err := r.Run(":8080")
	panic(err)
}

func registerRouters(r *gin.Engine, h *delivery.Handler) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/v1/user/register", h.Register)
	r.POST("/v1/user/login", h.Login)
}

func catRouters(r *gin.Engine, h *delivery.Handler) {
	r.Use(middleware.AuthMiddleware(&config.Config{}))
}
