package routes

import (
	"auth-service/handler"
	repository "auth-service/repository/user"
	"auth-service/service"

	"github.com/gin-gonic/gin"
)

func userGroupRoutes(r *gin.Engine) {
	userRepo := repository.NewUserRepositoryDB()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	testGruup := r.Group("api/v1/user")

	testGruup.POST("/create", userHandler.CreateUser)
	testGruup.GET("/", userHandler.GetUsers)
}
