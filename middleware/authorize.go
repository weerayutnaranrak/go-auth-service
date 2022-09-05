package middleware

import (
	"auth-service/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := ""
		if strings.Contains(authHeader, BEARER_SCHEMA) {
			tokenString = authHeader[len(BEARER_SCHEMA):]
		} else {
			tokenString = authHeader
		}
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
		} else {
			claims := token.Claims.(jwt.MapClaims)
			ctx.Set("user", claims)
		}

	}
}
