package main

import (
	"app/controllers"
	"app/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})

	models.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello GoPress!"})
	})

	r.POST("/v1/user/login", controllers.Login)

	r.Run(":" + os.Getenv("PORT"))
}
