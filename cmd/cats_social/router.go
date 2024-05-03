package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/delivery"
)

func initRouter(h *delivery.Handler, authMiddleware gin.HandlerFunc) {
	r := gin.Default()

	// registerRouters(app)
	registerRouters(r, h)
	catRouters(r, h, authMiddleware)
	matchRouters(r, h, authMiddleware)

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

func matchRouters(r *gin.Engine, h *delivery.Handler, authMiddleware gin.HandlerFunc) {
	r.Use(authMiddleware)

	r.POST("/v1/cat/match", h.MatchCat)
	r.GET("/v1/cat/match")
	r.POST("/v1/cat/match/approve")
	r.POST("/v1/cat/match/reject")
	r.DELETE("/v1/cat/match/:id", h.DeleteMatch)
}

func catRouters(r *gin.Engine, h *delivery.Handler, authMiddleware gin.HandlerFunc) {
	// example use case of authMiddleware
	r.Use(authMiddleware)
	r.POST("/v1/cat", h.RegisterCat)
}
