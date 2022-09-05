package routes

import (
	"auth-service/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authRoutes(r *gin.Engine) {
	authGroup := r.Group("api/v1/auth")
	authGroup.POST("/login", middleware.AuthorizeJWT(), func(ctx *gin.Context) {})
	authGroup.POST("/register", func(ctx *gin.Context) {})
	authGroup.Use(middleware.AuthorizeJWT())
	{
		authGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Ok",
			})
		})
	}

}
