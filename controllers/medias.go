package controllers

import (
	"app/models"
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var (
	storageClient *storage.Client
)

func UploadMedia(c *gin.Context) {
	db := models.DB

	bucket_name := os.Getenv("bucket_name")
	bucket_address := os.Getenv("bucket_address")
	google_credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	var err error

	ctx := appengine.NewContext(c.Request)

	storageClient, err := storage.NewClient(ctx,
		option.WithCredentialsFile(google_credentials),
	)
	if err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	defer f.Close()

	uploadedFile.Filename = slug.Make(uploadedFile.Filename)

	sw := storageClient.Bucket(bucket_name).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	if err := sw.Close(); err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	u, err := url.Parse("/" + bucket_name + "/" + sw.Attrs().Name)
	if err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	// save on database

	var media models.Media

	media.FileName = sw.Attrs().Name
	media.FileSrc = fmt.Sprintf("%s%s", bucket_address, u.EscapedPath())
	media.FileSize = sw.Attrs().Size
	media.MimeType = sw.Attrs().ContentType
	media.Title = sw.Attrs().Name

	if err := db.Create(&media).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	render(c, gin.H{"media": media})
}

func ListMedia(c *gin.Context) {
	db := models.DB
	var medias []models.Media

	db = db.Model(&models.Media{})

	// filter by title
	if title := c.Query("title"); title != "" {
		title = fmt.Sprintf("%%%s%%", title)
		db = db.Where("title ilike ?", title)
	}

	// pagination
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	p := CalculateOffset(page, pageSize)
	db = db.Limit(p.PageSize).Offset(p.Offset)

	if err := db.Find(&medias).Error; err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	renderList(c, gin.H{"medias": medias}, p.Page, p.Offset)
}

func GetMedia(c *gin.Context) {
	db := models.DB
	id := c.Param("id")
	var media models.Media

	db = db.Model(&models.Media{})

	if err := db.Find(&media, "id = ?", id).Error; err != nil {
		renderError(c, http.StatusNotFound, err)
		return
	}

	render(c, gin.H{"media": media})
}

func DeleteMedia(c *gin.Context) {
	db := models.DB
	var media models.Media
	db = db.Model(&models.Media{})
	id := c.Param("id")
	var err error

	bucket_name := os.Getenv("bucket_name")
	// bucket_address := os.Getenv("bucket_address")
	google_credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	ctx := appengine.NewContext(c.Request)

	storageClient, err := storage.NewClient(ctx,
		option.WithCredentialsFile(google_credentials),
	)
	if err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	// get data from database
	if err = db.Find(&media, "id = ?", id).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	// get object name
	object_name := media.FileName

	// if success, delete object from bucket
	obj := storageClient.Bucket(bucket_name).Object(object_name)
	log.Printf("object: %s", obj)
	attrs, err := obj.Attrs(ctx)
	log.Printf("Attrs: %s", attrs)

	if err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	obj = obj.If(storage.Conditions{GenerationMatch: attrs.Generation})

	log.Printf("object: %s", obj)

	if err = obj.Delete(ctx); err != nil {
		renderError(c, http.StatusBadRequest, err)
		return
	}

	// delete from database
	if err = db.Delete(&media, "id = ?", id).Error; err != nil {
		renderError(c, http.StatusInternalServerError, err)
		return
	}

	// return empty object
	render(c, gin.H{})

}
