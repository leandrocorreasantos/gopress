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

	if err := db.Where("user_id = ?", user_id).Find(&socialmediaprofiles).Error; err != nil {
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
	if err := db.Where(conditions, user_id, social_media_id).Find(&socialmediaprofile).Error; err != nil {
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

	if err := db.Create(&socialmediaprofile).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}

func UpdateSocialMediaProfile(c *gin.Context) {
	var socialmediaprofile models.SocialMediaProfile
	db := models.DB
	user_id := c.Param("user_id")
	social_media_id := c.Param("social_media_id")

	conditions := "user_id = ? and social_media_id = ?"
	if err := db.Where(conditions, user_id, social_media_id).Find(&socialmediaprofile).Error; err != nil {
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
	user_id := c.Param("user_id")
	social_media_id := c.Param("social_media_id")

	conditions := "user_id = ? and social_media_id = ?"
	if err := db.Where(conditions, user_id, social_media_id).Find(&socialmediaprofile).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := db.Where(conditions, user_id, social_media_id).Delete(&socialmediaprofile).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}
