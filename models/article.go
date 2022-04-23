package models

import (
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	ID       uint       `json:"id"`
	Name     string     `gorm:"size:100; not null" json:"name"`
	Articles []*Article `gorm:"many2many:articles_tags" json:"articles,omitempty"`
}

type MetaTag struct {
	ID   uint   `json:"id"`
	Name string `gorm:"size:100" json:"name"`
}

func (MetaTag) TableName() string {
	return "meta_tags"
}

type ArticleMetaTag struct {
	ID        uint     `json:"id"`
	MetaTagID uint     `json:"meta_tag_id"`
	ArticleID *uint    `json:"article_id,omitempty"`
	MetaTag   MetaTag  `json:"meta_tag"`
	Article   *Article `json:"article,omitempty"`
	TagValue  string   `gorm:"size:100;not null" json:"tag_value"`
}

func (ArticleMetaTag) TableName() string {
	return "articles_meta_tags"
}

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null;unique" json:"name,omitempty"`
	Slug string `gorm:"size:100;not null;unique" json:"slug,omitempty"`
}

func (c *Category) BeforeSave() error {
	c.Slug = slug.Make(c.Name)
	return nil
}

func (c *Category) FindIdBySlug(slug string) (id interface{}, err error) {
	db := DB
	if err := db.Find(&c, "slug = ?", slug).Error; err != nil {
		return nil, err
	}
	return c.ID, nil
}

func (c *Category) FindSlugByID(id uint) (slug interface{}, err error) {
	db := DB
	if err := db.Find(&c, "ID= ?", id).Error; err != nil {
		return nil, err
	}

	return c.Slug, nil
}

type Article struct {
	gorm.Model
	CategoryID    *uint             `json:"category_id,omitempty"`
	Category      *Category         `json:"category,omitempty"`
	UserID        *uint             `json:"user_id,omitempty"`
	User          *User             `json:"user,omitempty"`
	Title         string            `gorm:"size:255;not null;unique" json:"title,omitempty"`
	Slug          string            `gorm:"size:255;not null;unique" json:"slug,omitempty"`
	Subtitle      string            `gorm:"size:255" json:"subtitle,omitempty"`
	Content       string            `json:"content,omitempty"`
	IsDraft       bool              `gorm:"default:true" json:"is_draft,omitempty"`
	IsPublished   bool              `gorm:"default:false" json:"is_published,omitempty"`
	DatePublished time.Time         `json:"date_published"`
	Tags          []*Tag            `gorm:"many2many:articles_tags" json:"tags,omitempty"`
	MetaTags      []*ArticleMetaTag `json:"meta_tags,omitempty"`
	MediaID       uint              `json:"media_id,omitempty"`
	Media         Media             `json:"media,omitempty"`
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

func (a *Article) FindIdBySlug(slug string) (id interface{}, err error) {
	db := DB
	if err := db.Find(&a, "slug = ?", slug).Error; err != nil {
		return nil, err
	}
	return a.ID, nil
}

func (a *Article) FindSlugByID(id uint) (slug interface{}, err error) {
	db := DB
	if err := db.Find(&a, "ID= ?", id).Error; err != nil {
		return nil, err
	}

	return a.Slug, nil
}
