package controllers

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategories(c *gin.Context) {
	var categories []models.Category
	db := models.DB

	// filter by name
	if name := c.DefaultQuery("name", ""); name != "" {
		name = fmt.Sprintf("%%%s%%", name)
		db = db.Where("name ilike ?", name)
	}

	// find categories
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

	db = db.Model(&models.Category{})

	if err := db.First(&category, "ID = ?", id).Error; err != nil {
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

	renderCreate(c, gin.H{})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	db := models.DB
	id := c.Param("id")

	if err := db.Where("ID = ?", id).First(&category).Error; err != nil {
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

	db = db.Model(&models.Category{})

	if err := db.First(&category, "ID= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := db.Where("ID= ?", id).Delete(&category).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}
