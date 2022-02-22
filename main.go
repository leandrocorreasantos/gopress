package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main(){
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "Hello GoPress!"})
	})

	r.Run(":" + os.Getenv("PORT"))
}
