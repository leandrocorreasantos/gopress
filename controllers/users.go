package controllers

import (
	"app/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/joho/godotenv/autoload"
)

func Login(c *gin.Context) {
	var user models.User
	var login models.Login

	db := models.DB

	if err := c.ShouldBindJSON(&login); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	err := db.Where("username = ?", login.Username).Find(&user).Error
	if err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	err = models.VerifyPassword(user.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	// expiration time
	iat := time.Now().Unix()
	exp := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"iat":      iat,
		"exp":      exp,
	})

	secretKey := os.Getenv("secretKey")
	tokenString, e := token.SignedString([]byte(secretKey))
	if e != nil {
		renderError(c, http.StatusBadRequest, e)
		return
	}

	render(c, gin.H{"access_token": tokenString})
}

func FindUser(c *gin.Context) {
	render(c, gin.H{"user": "Hi, user!"})
}
