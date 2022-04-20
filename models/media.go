package models

import (
	"github.com/jinzhu/gorm"
)

type Media struct {
	gorm.Model
	FileName string `gorm:"size:255;not null" json:"file_name"`
	FileSrc  string `gorm:"size:255;not null" json:"file_src"`
	FileSize int64  `gorm:"default:0" json:"file_size"`
	MimeType string `gorm:"size:30;" json:"mime_type"`
	Title    string `gorm:"size:255;not null" json:"title"`
	AltText  string `gorm:"size:255" json:"alt_text"`
	Copyight string `gorm:"size:255" json:"copyright"`
}
