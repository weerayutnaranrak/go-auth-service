package handler

import (
	"auth-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{
		userSrv: userSrv,
	}
}

func (h userHandler) CreateUser(ctx *gin.Context) {
	user, _ := h.userSrv.NewUser(service.NewUserRequest{Name: "test2"})
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
