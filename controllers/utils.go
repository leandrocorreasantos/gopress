package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(3)
}

func render(c *gin.Context, data gin.H) {
	status := http.StatusOK
	response := gin.H{
		"status": status,
		"data":   data,
	}
	c.JSON(status, response)
}

func renderList(c *gin.Context, data gin.H, page int, offset int, total int) {
	status := http.StatusOK

	response := gin.H{
		"status": status,
		"page":   page,
		"offset": offset,
		"total":  total,
		"data":   data,
	}
	c.JSON(status, response)
}

func renderError(c *gin.Context, status int, err error) {
	response := gin.H{"code": status, "error": err}
	c.JSON(status, response)
}
