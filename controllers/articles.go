package controllers

import (
	"app/models"
	"fmt"
	// "log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// admin functions
func ListArticles(c *gin.Context) {
	time_layout := "2006-01-02"
	var articles []models.Article
	var category models.Category
	var user models.User
	db := models.DB

	var err error

	db = db.Model(&models.Article{})
	// include other models
	db = db.Preload("User").Preload("Category").Preload("Tags")
	db = db.Preload("MetaTags").Preload("MetaTags.MetaTag").Preload("Media")

	// filter by date_published
	date_published_end, err := time.Parse(
		time_layout,
		c.Query("date_published_end"),
	)
	if err != nil {
		date_published_end = time.Now()
	}
	date_published_start, err := time.Parse(
		time_layout,
		c.Query("date_published_start"),
	)
	if err == nil {
		db = db.Where(
			"date_published between ? and ?",
			date_published_start, date_published_end,
		)
	}

	// Category
	if category_slug := c.Query("category"); category_slug != "" {
		category_id, err := category.FindIdBySlug(category_slug)
		if err != nil {
			renderError(c, http.StatusNotFound, err)
			return
		} else {
			db = db.Where("category_id = ?", category_id)
		}
	}

	// User (Author)
	if username := c.Query("author"); username != "" {
		user_id, err := user.FindIdByUsername(username)
		if err != nil {
			renderError(c, http.StatusNotFound, err)
			return
		} else {
			db = db.Where("user_id = ?", user_id)
		}
	}

	// filter by is_draft
	draft := c.DefaultQuery("is_draft", "true")
	if is_draft, err := strconv.ParseBool(draft); err == nil {
		db = db.Where("is_draft = ?", is_draft)
	}
	published := c.DefaultQuery("is_published", "false")
	if is_published, err := strconv.ParseBool(published); err == nil {
		db = db.Where("is_published = ?", is_published)
	}

	// filter by title
	if title := c.Query("title"); title != "" {
		title = fmt.Sprintf("%%%s%%", title)
		db = db.Where("title ilike ?", title)
	}

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	p := CalculateOffset(page, pageSize)
	db = db.Limit(p.PageSize).Offset(p.Offset)

	// find results
	if err := db.Find(&articles).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	renderList(c, gin.H{"articles": articles}, p.Page, p.Offset)
}

func GetArticle(c *gin.Context) {
	var article models.Article
	db := models.DB
	id := c.Param("id")

	db = db.Model(&models.Article{})
	db = db.Preload("User").Preload("Category").Preload("Tags")
	db = db.Preload("MetaTags").Preload("MetaTags.MetaTag").Preload("Media")

	if err := db.Where("ID= ?", id).Find(&article).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"article": article})
}

func CreateArticle(c *gin.Context) {
	var article models.Article
	db := models.DB

	if err := c.ShouldBindJSON(&article); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Create(&article).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	renderCreate(c, gin.H{})
}

func UpdateArticle(c *gin.Context) {
	var article models.Article
	db := models.DB
	id := c.Param("id")

	db = db.Model(&models.Article{})

	if err := db.First(&article, "id= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := c.ShouldBindJSON(&article); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	if err := db.Save(&article).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	render(c, gin.H{})
}

func DeleteArticle(c *gin.Context) {
	db := models.DB
	id := c.Param("id")
	var article models.Article

	db = db.Model(&models.Article{})

	if err := db.First(&article, "id= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	if err := db.Where("id= ?", id).Delete(&article).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}

func DraftArticle(c *gin.Context) {
	db := models.DB
	id := c.Param("id")
	var article models.Article

	db = db.Model(&models.Article{})

	if err := db.First(&article, "ID= ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	article.MarkAsDraft()

	if err := db.Save(&article).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}

func PublishArticle(c *gin.Context) {
	db := models.DB
	id := c.Param("id")
	var article models.Article

	db = db.Model(&models.Article{})

	if err := db.First(&article, "ID = ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	article.MarkAsPublished()

	if err := db.Save(&article).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{})
}

// site functions

func ListPublishedArticles(c *gin.Context) {
	time_layout := "2006-01-02"
	var articles []models.Article
	var category models.Category
	var user models.User
	db := models.DB

	db = db.Model(&models.Article{})

	db = db.Preload("User").Preload("Category").Preload("Tags")
	db = db.Preload("MetaTags").Preload("MetaTags.MetaTag").Preload("Media")

	// filter by date_published
	date_published_end, err := time.Parse(
		time_layout,
		c.Query("date_published_end"),
	)
	if err != nil {
		date_published_end = time.Now()
	}
	date_published_start, err := time.Parse(
		time_layout,
		c.Query("date_published_start"),
	)
	if err == nil {
		db = db.Where(
			"date_published between ? and ?",
			date_published_start, date_published_end,
		)
	}

	// Category
	if category_slug := c.Query("category"); category_slug != "" {
		category_id, err := category.FindIdBySlug(category_slug)
		if err != nil {
			renderError(c, http.StatusNotFound, err)
			return
		} else {
			db = db.Where("category_id = ?", category_id)
		}
	}

	// User (Author)
	if username := c.Query("author"); username != "" {
		user_id, err := user.FindIdByUsername(username)
		if err != nil {
			renderError(c, http.StatusNotFound, err)
			return
		} else {
			db = db.Where("user_id = ?", user_id)
		}
	}

	// filter by title
	if title := c.Query("title"); title != "" {
		title = fmt.Sprintf("%%%s%%", title)
		db = db.Where("title ilike ?", title)
	}

	// only publisheds:
	db = db.Where("is_published = true")

	// add pagination
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	p := CalculateOffset(page, pageSize)
	db = db.Limit(p.PageSize).Offset(p.Offset)

	// find results
	if err := db.Find(&articles).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	renderList(c, gin.H{"articles": articles}, p.Page, p.Offset)
}

func GetPublishedArticle(c *gin.Context) {
	// get by slug
	var article models.Article
	db := models.DB
	slug := c.Param("slug")

	db = db.Model(&models.Article{})

	db = db.Preload("User").Preload("Category").Preload("Tags")
	db = db.Preload("MetaTags").Preload("MetaTags.MetaTag").Preload("Media")

	if err := db.First(&article, "Slug= ?", slug).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"article": article})
}
