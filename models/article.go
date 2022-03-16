package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null;unique" json:"name"`
	Slug string `gorm:"size:100;not null;" json:"slug"`
}

type Article struct {
	gorm.Model
	CategoryID    uint      `json:"category_id"`
	Category      Category  `json:"category"`
	UserID        uint      `json:"user_id"`
	User          User      `json:"user"`
	Title         string    `gorm:"size:255;not null;unique" json:"title"`
	Slug          string    `gorm:"size:255;not null;" json:"slug"`
	Subtitle      string    `gorm:"size:255" json:"subtitle"`
	Content       string    `json:"content"`
	IsDraft       bool      `gorm:"default:true" json:"is_draft"`
	IsPublished   bool      `gorm:"default:false" json:"is_published"`
	DatePublished time.Time `json:"date_published"`
}

func (a *Article) MarkAsPublished() error {
	a.IsPublished = true
	a.IsDraft = false
	a.DatePublished = time.Now()
	return nil
}

func (a *Article) MarkAsDraft() error {
	a.IsPublished = false
	a.IsDraft = true
	return nil
}
