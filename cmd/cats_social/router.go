package main

import "github.com/gin-gonic/gin"

func initRouter() {
	r := gin.Default()

	// registerRouters(app)
	registerRouters(r)

	err := r.Run(":8080")
	panic(err)
}

func registerRouters(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
