package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListMetaTag(c *gin.Context) {
	var meta_tags []models.MetaTag
	db := models.DB

	if err := db.Find(&meta_tags).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{"meta_tags": meta_tags})
}

func GetMetaTag(c *gin.Context) {
	var meta_tag models.MetaTag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&meta_tag, "id = ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"meta_tag": meta_tag})
}

func CreateMetaTag(c *gin.Context) {
	var meta_tag models.MetaTag
	db := models.DB

	if err := c.ShouldBindJSON(&meta_tag); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Create(&meta_tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	renderCreate(c, gin.H{})
}

func UpdateMetaTag(c *gin.Context) {
	var meta_tag models.MetaTag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&meta_tag, "ID= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := c.ShouldBindJSON(&meta_tag); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Save(&meta_tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}

func DeleteMetaTag(c *gin.Context) {
	var meta_tag models.MetaTag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&meta_tag, "ID= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := db.Where("id = ?", id).Delete(&meta_tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}
