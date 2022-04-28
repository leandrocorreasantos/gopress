package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSocialMedia(c *gin.Context) {
	var socialmedias []models.SocialMedia

	db := models.DB

	if err := db.Find(&socialmedias).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{"social_medias": socialmedias})
}

func GetSocialMedia(c *gin.Context) {
	var social_media models.SocialMedia
	db := models.DB
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&social_media).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"social_media": social_media})
}

func CreateSocialMedia(c *gin.Context) {
	db := models.DB
	var socialmedia models.SocialMedia

	if err := c.ShouldBindJSON(&socialmedia); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Create(&socialmedia).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	renderCreate(c, gin.H{})
}

func UpdateSocialMedia(c *gin.Context) {
	db := models.DB
	var socialmedia models.SocialMedia
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&socialmedia).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&socialmedia); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Save(&socialmedia).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{})
}

func DeleteSocialMedia(c *gin.Context) {
	db := models.DB
	var socialmedia models.SocialMedia
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&socialmedia).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := db.Where("id = ?", id).Delete(&socialmedia).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}
