package middlewares

import (
	"CRUD-GIN/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error: " : "Auth header is required",
			})
			ctx.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// No Bearer found
			ctx.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid token format"})
			return
		}

		token, err := utils.ValidateJWT(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid or expired token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Could not parse token claims"})
			ctx.Abort()
			return
		}

		_, ok = claims["username"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in claims"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
