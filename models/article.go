package models

import (
	"time"

	"github.com/gosimple/slug"

	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null;unique" json:"name"`
	Slug string `gorm:"size:100;not null;unique" json:"slug"`
}

func (c *Category) BeforeSave() error {
	c.Slug = slug.Make(c.Name)
	return nil
}

type Article struct {
	gorm.Model
	CategoryID    uint      `json:"category_id"`
	Category      Category  `json:"category"`
	UserID        uint      `json:"user_id"`
	User          User      `json:"user"`
	Title         string    `gorm:"size:255;not null;unique" json:"title"`
	Slug          string    `gorm:"size:255;not null;unique" json:"slug"`
	Subtitle      string    `gorm:"size:255" json:"subtitle"`
	Content       string    `json:"content"`
	IsDraft       bool      `gorm:"default:true" json:"is_draft"`
	IsPublished   bool      `gorm:"default:false" json:"is_published"`
	DatePublished time.Time `json:"date_published"`
}

func (a *Article) BeforeSave() error {
	// create slug
	a.Slug = slug.Make(a.Title)
	// alter published status
	if a.IsPublished == true {
		a.MarkAsPublished()
	} else {
		a.MarkAsDraft()
	}
	return nil
}

func (a *Article) MarkAsPublished() error {
	// update data to mark as published
	a.IsPublished = true
	a.IsDraft = false
	a.DatePublished = time.Now()
	return nil
}

func (a *Article) MarkAsDraft() error {
	// update data to mark as draft
	a.IsPublished = false
	a.IsDraft = true
	return nil
}
