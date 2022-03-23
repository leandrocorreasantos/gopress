package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB

func ConnectDB() {
	dbuser := os.Getenv("databaseUser")
	dbpass := os.Getenv("databasePass")
	dbname := os.Getenv("databaesName")
	dbhost := os.Getenv("databaseHost")
	dbdriver := os.Getenv("databaseDriver")

	log.Printf("%s://%s:%s@%s/%s?sslmode=disable",
		dbdriver, dbuser, dbpass, dbhost, dbname)

	dbURI := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		dbdriver,
		dbuser,
		dbpass,
		dbhost,
		dbname,
	)

	db, err := gorm.Open(dbdriver, dbURI)

	if err != nil {
		log.Fatal("Error while connecting to database ", err)
		return
	}

	// drop enum type if exists
	_ = db.Exec("DROP TYPE IF EXISTS UserRole;")
	// create enum types
	_ = db.Exec("CREATE TYPE UserRole as ENUM ('admin', 'editor', 'author', 'reader')")

	// migrations
	db.AutoMigrate(&User{})
	db.AutoMigrate(&SocialMedia{})
	db.AutoMigrate(&SocialMediaProfile{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Article{})

	// create super User
	var superuser User
	var count int64
	username := os.Getenv("superuserUsername")
	password := os.Getenv("superuserPassword")
	db.Model(&superuser).Where("username = ?", username).Count(&count)
	log.Println("number of superusers: ", count)
	if count == 0 {
		superuser.ID = 1
		superuser.Username = username
		superuser.Password = password
		superuser.Active = true
		superuser.Role = "admin"
		superuser.Birthday = "1970-01-01"
		if err := db.Create(&superuser).Error; err != nil {
			log.Println("Error creating super user: ", err.Error())
			return
		}
	}

	log.Println("Database connected")

	DB = db

}
