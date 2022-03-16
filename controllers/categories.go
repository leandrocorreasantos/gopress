package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategories(c *gin.Context) {
	var categories []models.Category

	db := models.DB

	if err := db.Find(&categories).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{"categories": categories})
}

func GetCategory(c *gin.Context) {
	var category models.Category

	db := models.DB
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"category": category})
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	db := models.DB

	if err := c.ShouldBindJSON(&category); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Create(&category).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	db := models.DB
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Save(&category).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{})
}

func DeleteCategory(c *gin.Context) {
	db := models.DB
	var category models.Category
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := db.Where("id = ?", id).Delete(&category).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}
