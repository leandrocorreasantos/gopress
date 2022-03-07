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
		// user
		admin.GET("/user", controllers.ShowUser)    // logged user
		admin.GET("/users", controllers.ListUsers)  // all users
		admin.GET("/user/:id", controllers.GetUser) // one user
		admin.POST("/user", controllers.CreateUser)
		admin.PUT("/user/:id", controllers.UpdateUser)
		admin.PUT("/user/:id/password", controllers.UpdatePassword)
		// social media
		admin.GET("/socialmedia", controllers.ListSocialMedia)
		admin.GET("/socialmedia/:id", controllers.GetSocialMedia)
		admin.POST("/socialmedia", controllers.CreateSocialMedia)
		admin.PUT("/socialmedia/:id", controllers.UpdateSocialMedia)
		admin.DELETE("/socialmedia/:id", controllers.DeleteSocialMedia)
	}

	r.Run(":" + os.Getenv("PORT"))
}
