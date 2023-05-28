package helper

import (
	"net/http"
	"pancakaki/internal/domain/web"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtKey string) gin.HandlerFunc {
	var jwtKeyByte = []byte(jwtKey)

	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return jwtKeyByte, nil
		})
		if !token.Valid || err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)

		ctx.Next()
	}
}

func AuthCustomer(context *gin.Context) (*gin.Context, interface{}) {
	claims := context.MustGet("claims").(jwt.MapClaims)
	idClaim := claims["id"]
	role := claims["role"].(string)
	if role != "customer" {
		result := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  "UNAUTHORIZED",
			Message: "unauthorized",
			Data:    "user is unauthorized",
		}
		context.JSON(http.StatusUnauthorized, result)
		return context, 0
	}
	return context, idClaim
}
