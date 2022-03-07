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

	admin := r.Group("/v1/admin")
	admin.Use(models.JwtVerify())
	{
		admin.GET("/user", controllers.ShowUser)
		admin.GET("/users", controllers.ListUsers)
		admin.POST("/user", controllers.CreateUser)
		admin.PUT("/user/:id", controllers.UpdateUser)
		admin.PUT("/user/:id/password", controllers.UpdatePassword)
	}

	r.Run(":" + os.Getenv("PORT"))
}
