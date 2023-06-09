package middleware

import (
	"fmt"
	"github.com/ecom/pkg/controller"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT -> to authorize JWT Token
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})

		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := controller.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)

			} else {
				if token.Valid {
					ctx.Set("userID", claims["userID"])
					fmt.Println("during authorization", claims["userID"])
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}
			}
		}
	}
}