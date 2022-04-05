package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListArticleMetaTag(c *gin.Context) {
	var article_meta_tags []models.ArticleMetaTag
	db := models.DB

	if err := db.Find(&article_meta_tags).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{"article_meta_tags": article_meta_tags})
}

func GetArticleMetaTag(c *gin.Context) {
	var article_meta_tag models.ArticleMetaTag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&article_meta_tag, "id = ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"article_meta_tag": article_meta_tag})
}

func CreateArticleMetaTag(c *gin.Context) {
	var article_meta_tag models.ArticleMetaTag
	db := models.DB

	if err := c.ShouldBindJSON(&article_meta_tag); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Create(&article_meta_tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	renderCreate(c, gin.H{})
}

func UpdateArticleMetaTag(c *gin.Context) {
	var article_meta_tag models.ArticleMetaTag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&article_meta_tag, "ID= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := c.ShouldBindJSON(&article_meta_tag); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Save(&article_meta_tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}

func DeleteArticleMetaTag(c *gin.Context) {
	var article_meta_tag models.ArticleMetaTag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&article_meta_tag, "ID= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := db.Where("id = ?", id).Delete(&article_meta_tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}
