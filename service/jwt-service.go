package service

import (
	"fmt"
	"goTest/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var secretKey = []byte(fmt.Sprintf("%s", viper.Get("JWT.secret")))

type MyCustomClaims struct {
	Name  string `json:"foo"`
	Email string
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) string {

	claims := MyCustomClaims{
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(secretKey)

	return ss
}

func AuthToken(ctx *gin.Context) error {
	tokenString := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		if claims.RegisteredClaims.Issuer == "test" {
			return nil
		}
	}

	return err
}
