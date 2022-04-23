package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(3)
}

type Paginate struct {
	Offset   int
	PageSize int
	Page     int
}

func CalculateOffset(page, pageSize int) Paginate {
	if page == 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	return Paginate{offset, pageSize, page}
}

// rendering results
func render(c *gin.Context, data gin.H) {
	status := http.StatusOK
	response := gin.H{
		"status": status,
		"data":   data,
	}
	c.JSON(status, response)
}

func renderCreate(c *gin.Context, data gin.H) {
	status := http.StatusCreated
	response := gin.H{
		"status": status,
		"data":   data,
	}
	c.JSON(status, response)
}

func renderList(c *gin.Context, data gin.H, page int, offset int) {
	status := http.StatusOK

	response := gin.H{
		"status": status,
		"page":   page,
		"offset": offset,
		"data":   data,
	}
	c.JSON(status, response)
}

func renderError(c *gin.Context, status int, err error) {
    log.Printf("Erro %d: %s", status, err.Error())
	response := gin.H{"code": status, "error": err.Error()}
	c.JSON(status, response)
}
