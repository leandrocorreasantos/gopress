package controllers

import (
	"app/models"
	"fmt"
	"net/http"
	"os"
	"strings"
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
        aud := os.Getenv("JWTAudience")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"iat":      iat,
		"exp":      exp,
                "aud":      aud,
	})

	secretKey := os.Getenv("secretKey")
	tokenString, e := token.SignedString([]byte(secretKey))
	if e != nil {
		renderError(c, http.StatusBadRequest, e)
		return
	}

	render(c, gin.H{"access_token": tokenString})
}

func ShowUser(c *gin.Context) {
	var user models.User
	secretKey := os.Getenv("secretKey")
	db := models.DB

	tk := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

	tokenauth := fmt.Sprintf("%s", tk)

	token, err := jwt.ParseWithClaims(tokenauth, &models.Token{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		renderError(c, http.StatusUnauthorized, err)
		return
	}

	claims, ok := token.Claims.(*models.Token)
	if !ok && token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
		return
	}

	db = db.Model(&models.User{}).Preload("SocialMediaProfile")
	db = db.Preload("SocialMediaProfile.SocialMedia")

	if err := db.First(&user, "username = ?", claims.Username).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{"user": user})
}

func GetUser(c *gin.Context) {
	db := models.DB
	var user models.User
	id := c.Param("id")

	db = db.Model(&models.User{}).Preload("SocialMediaProfile")
	db = db.Preload("SocialMediaProfile.SocialMedia")

	if err := db.Where("ID = ?", id).First(&user).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"user": user})
}

func ListUsers(c *gin.Context) {
	db := models.DB
	var users []models.User

	db = db.Model(&models.User{}).Preload("SocialMediaProfile")
	db = db.Preload("SocialMediaProfile.SocialMedia")

	err := db.Find(&users).Error
	if err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}
	render(c, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
	db := models.DB
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	renderCreate(c, gin.H{})
}

func UpdateUser(c *gin.Context) {
	db := models.DB
	var user models.User
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Omit("password").Save(&user).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{})
}

func UpdatePassword(c *gin.Context) {
	db := models.DB
	var user models.User
	var input models.ChangePassword
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	errf := models.VerifyPassword(user.Password, input.OldPassword)
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		renderError(c, http.StatusBadRequest, errf)
		return
	}

	user.Password = input.NewPassword
	if err := db.Save(&user).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{})
}
