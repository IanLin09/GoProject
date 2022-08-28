package middleware

import (
	"goTest/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		err := service.AuthToken(ctx)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Invalid Token.",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
