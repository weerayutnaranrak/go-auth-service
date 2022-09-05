package routes

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")

	authGroup.POST("/login", func(ctx *gin.Context) {})
	authGroup.POST("/register", func(ctx *gin.Context) {})
}
