package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/app/delivery"
)

func initRouter(h *delivery.Handler, authMiddleware gin.HandlerFunc) {
	r := gin.Default()

	// registerRouters(app)
	registerRouters(r, h)
	registerCatRouters(r, h, authMiddleware)
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
	r.GET("/v1/cat/match", h.GetMatch)
	r.POST("/v1/cat/match/approve", h.ApproveMatch)
	r.POST("/v1/cat/match/reject", h.RejectMatch)
	r.DELETE("/v1/cat/match/:id", h.DeleteMatch)
}

func registerCatRouters(r *gin.Engine, h *delivery.Handler, authMiddleware gin.HandlerFunc) {
	// example use case of authMiddleware
	r.Use(authMiddleware)
	r.POST("/v1/cat", h.RegisterCat)
	r.GET("/v1/cat", h.GetCat)
	r.PUT("/v1/cat/:id", h.PutCat)
	r.DELETE("/v1/cat/:id", h.DeleteCat)
}
