package handler

import (
	userService "auth-service/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSrv userService.UserService
}

func NewUserHandler(userSrv userService.UserService) userHandler {
	return userHandler{
		userSrv: userSrv,
	}
}

func (h userHandler) CreateUser(ctx *gin.Context) {
	user, _ := h.userSrv.NewUser(userService.NewUserRequest{Name: "test2"})
	ctx.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
func (h userHandler) GetUsers(ctx *gin.Context) {
	users, _ := h.userSrv.GetUsers(1)
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
