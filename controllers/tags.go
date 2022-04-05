package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTag(c *gin.Context) {
	var tags []models.Tag
	db := models.DB

	if err := db.Find(&tags).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{"tags": tags})
}

func GetTag(c *gin.Context) {
	var tag models.Tag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&tag, "id = ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"tag": tag})
}

func CreateTag(c *gin.Context) {
	var tag models.Tag
	db := models.DB

	if err := c.ShouldBindJSON(&tag); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Create(&tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	renderCreate(c, gin.H{})
}

func UpdateTag(c *gin.Context) {
	var tag models.Tag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&tag, "ID= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := c.ShouldBindJSON(&tag); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Save(&tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}

func DeleteTag(c *gin.Context) {
	var tag models.Tag
	db := models.DB
	id := c.Param("id")

	if err := db.First(&tag, "ID= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := db.Where("id = ?", id).Delete(&tag).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}
