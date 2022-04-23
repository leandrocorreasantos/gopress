package main

import (
	"app/controllers"
	"app/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})

	models.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello GoPress!"})
	})

	r.POST("/v1/user/login", controllers.Login)

    // articles
    r.GET("/articles", controllers.ListPublishedArticles)
    r.GET("/articles/:slug", controllers.GetPublishedArticle)

	admin := r.Group("/v1/admin")
	admin.Use(models.JwtVerify())
	{
		// user
		admin.GET("/user", controllers.ShowUser)    // logged user
		admin.GET("/users", controllers.ListUsers)  // all users
		admin.GET("/user/:id", controllers.GetUser) // one user
		admin.POST("/user", controllers.CreateUser)
		admin.PUT("/user/:id", controllers.UpdateUser)
		admin.PUT("/user/:id/password", controllers.UpdatePassword)
		// social media
		admin.GET("/socialmedia", controllers.ListSocialMedia)
		admin.GET("/socialmedia/:id", controllers.GetSocialMedia)
		admin.POST("/socialmedia", controllers.CreateSocialMedia)
		admin.PUT("/socialmedia/:id", controllers.UpdateSocialMedia)
		admin.DELETE("/socialmedia/:id", controllers.DeleteSocialMedia)
		// social media profile
		admin.GET("/socialmediaprofile/:user_id", controllers.ListSocialMediaProfile)
		admin.GET("/socialmediaprofile/:user_id/:social_media_id", controllers.GetSocialMediaProfile)
		admin.POST("/socialmediaprofile", controllers.CreateSocialMediaProfile)
		admin.PUT("/socialmediaprofile/:id", controllers.UpdateSocialMediaProfile)
		admin.DELETE("/socialmediaprofile/:id", controllers.DeleteSocialMediaProfile)
		// category
		admin.GET("/category", controllers.ListCategories)
		admin.GET("/category/:id", controllers.GetCategory)
		admin.POST("/category", controllers.CreateCategory)
		admin.PUT("/category/:id", controllers.UpdateCategory)
		admin.DELETE("/category/:id", controllers.DeleteCategory)
		// articles
		admin.GET("/articles", controllers.ListArticles)
		admin.GET("/article/:id", controllers.GetArticle)
		admin.POST("/article", controllers.CreateArticle)
		admin.PUT("/article/:id", controllers.UpdateArticle)
		admin.PATCH("/article/:id/draft", controllers.DraftArticle)
		admin.PATCH("/article/:id/publish", controllers.PublishArticle)
		admin.DELETE("/article/:id", controllers.DeleteArticle)
		// tags
		admin.GET("/tag", controllers.ListTag)
		admin.GET("/tag/:id", controllers.GetTag)
		admin.POST("/tag", controllers.CreateTag)
		admin.PUT("/tag/:id", controllers.UpdateTag)
		admin.DELETE("/tag/:id", controllers.DeleteTag)
		// meta tags
		admin.GET("/meta_tag", controllers.ListMetaTag)
		admin.GET("/meta_tag/:id", controllers.GetMetaTag)
		admin.POST("/meta_tag", controllers.CreateMetaTag)
		admin.PUT("/meta_tag/:id", controllers.UpdateMetaTag)
		admin.DELETE("/meta_tag/:id", controllers.DeleteMetaTag)
		// article meta tags
		admin.GET("/article_meta_tag", controllers.ListArticleMetaTag)
		admin.GET("/article_meta_tag/:id", controllers.GetArticleMetaTag)
		admin.POST("/article_meta_tag", controllers.CreateArticleMetaTag)
		admin.PUT("/article_meta_tag/:id", controllers.UpdateArticleMetaTag)
		admin.DELETE("/article_meta_tag/:id", controllers.DeleteArticleMetaTag)
		// upload media
		admin.GET("/media", controllers.ListMedia)
		admin.GET("/media/:id", controllers.GetMedia)
		admin.POST("/media/upload", controllers.UploadMedia)
		admin.DELETE("/media/:id", controllers.DeleteMedia)
	}

	r.Run(":" + os.Getenv("PORT"))
}
