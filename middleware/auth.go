package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwTsecrets = []byte("secret-key_archit_1234_231")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("authToken")
		if err != nil || authCookie == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}
		tokenString := authCookie
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwTsecrets, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userId := uint(claims["user_id"].(float64))
			c.Set("user_id", userId)
			c.Set("role", claims["role"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
		c.Next()
	}

}
