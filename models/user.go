package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRole string

const (
	ADMIN  UserRole = "admin"
	EDITOR UserRole = "edior"
	AUTHOR UserRole = "author"
	READER UserRole = "reader"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"size:100;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Active    bool      `gorm:"default:true" json:"active"`
	FirstName string    `gorm:"size:100" json:"first_name"`
	LastName  string    `gorm:"size:100" json:"last_name"`
	Birthday  time.Time `sql:"timestamp without time zone" json:"birthday"`
	Biography string    `gorm:"type:text" json:"biography"`
	Role      UserRole  `gorm:"not null;default:'reader'" json:"role"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// create hashed password
func (u *User) BeforeSave() error {
	if u.Password != "" {
		hashedPassword, err := HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hadhedPass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hadhedPass), []byte(pass))
}
