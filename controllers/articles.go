package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// admin functions
func ListArticles(c *gin.Context) {
	var articles []models.Article
	db := models.DB
	// filter by: is_draft, is_published, date_published, Createdat, UpdatedAt
	// Category, User, Title

	// add pagination

	if err := db.Find(&articles).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{"articles": articles})
}

func GetArticle(c *gin.Context) {
	var article models.Article
	db := models.DB
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"article": article})
}

func CreateArticle(c *gin.Context) {}

func UpdateArticle(c *gin.Context) {}

func DeleteArticle(c *gin.Context) {}

func ToDraftArticle(c *gin.Context) {}

func PublishArticle(c *gin.Context) {}

// site functions

func ListPublishedArticles(c *gin.Context) {
	// filter by date_published, category, user, Title
	// add pagination

}

func GetPublishedArticle(c *gin.Context) {}

func GetNeighborArticles(c *gin.Context) {
	// id := c.Param("id")
	// get the article before and after
}
