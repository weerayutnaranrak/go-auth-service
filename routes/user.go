package routes

import (
	"auth-service/handler"
	userRepository "auth-service/repository/user"
	userService "auth-service/service/user"

	"github.com/gin-gonic/gin"
)

func userGroupRoutes(r *gin.Engine) {
	userRepo := userRepository.NewUserRepositoryDB()
	userService := userService.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	testGruup := r.Group("api/v1/user")

	testGruup.POST("/create", userHandler.CreateUser)
	testGruup.GET("/", userHandler.GetUsers)
}
