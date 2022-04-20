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
	"net/http"
	"net/url"
	"os"
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

	// render(c, gin.H{"filename": sw.Attrs().Name, "filesrc": u.EscapedPath()})
}
