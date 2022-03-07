package models

type SocialMedia struct {
	ID       uint   `json:"id"`
	Name     string `gorm:"size:100;" json:"name"`
	BaseUrl  string `gorm:"size:255;not null" json:"base_url"`
	UserChar string `gorm:"size:1;default:'@';not null" json:"user_char"`
}

type SocialMediaProfile struct {
	ID            uint        `json:"id"`
	User          User        `json:"user"`
	UserID        uint        `json:"user_id"`
	SocialMedia   SocialMedia `json:"social_media"`
	SocialMediaID uint        `json:"social_media_id"`
	Nickname      string      `gorm:"size:100;not null" json:"nickname"`
}
