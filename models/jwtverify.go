package models

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("secretKey")

		tk := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		if tk == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token empty"})
			return
		}

		tokenauth := fmt.Sprintf("%s", tk)
		if tokenauth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenauth, &Token{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			},
		)

		if err != nil {
			log.Println("Parse with claims error: ", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error})
			return
		}

		if _, ok := token.Claims.(*Token); !ok && !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			return
		}

		c.Next()
	}
}
