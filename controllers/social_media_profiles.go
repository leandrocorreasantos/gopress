package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSocialMediaProfile(c *gin.Context) {
	var socialmediaprofiles []models.SocialMediaProfile
	db := models.DB
	user_id := c.Param("user_id")

	db = db.Model(&models.SocialMediaProfile{})

	db = db.Where("user_id = ?", user_id).Preload("SocialMedia").Preload("User")
	if err := db.Find(&socialmediaprofiles).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"social_media_profiles": socialmediaprofiles})
}

func GetSocialMediaProfile(c *gin.Context) {
	var socialmediaprofile models.SocialMediaProfile
	db := models.DB
	user_id := c.Param("user_id")
	social_media_id := c.Param("social_media_id")

	conditions := "user_id = ? and social_media_id = ?"
	db = db.Preload("User").Preload("SocialMedia")
	db = db.Where(conditions, user_id, social_media_id)
	if err := db.Find(&socialmediaprofile).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"social_media_profile": socialmediaprofile})
}

func CreateSocialMediaProfile(c *gin.Context) {
	var socialmediaprofile models.SocialMediaProfile
	db := models.DB

	if err := c.ShouldBindJSON(&socialmediaprofile); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Save(&socialmediaprofile).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	renderCreate(c, gin.H{})
}

func UpdateSocialMediaProfile(c *gin.Context) {
	var socialmediaprofile models.SocialMediaProfile
	db := models.DB
	id := c.Param("id")

	if err := db.Find(&socialmediaprofile, "id = ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := c.ShouldBindJSON(&socialmediaprofile); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Save(&socialmediaprofile).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{})
}

func DeleteSocialMediaProfile(c *gin.Context) {
	var socialmediaprofile models.SocialMediaProfile
	db := models.DB
	id := c.Param("id")

	if err := db.Find(&socialmediaprofile, "id = ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	db = db.Where("id = ?", id)
	if err := db.Delete(&socialmediaprofile).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}
