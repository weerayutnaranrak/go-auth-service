package routes

import "github.com/gin-gonic/gin"

func Serve(r *gin.Engine) {
	userGroupRoutes(r)
	authRoutes(r)
}
