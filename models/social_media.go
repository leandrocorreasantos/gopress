package models

type SocialMedia struct {
	ID       uint    `json:"id"`
	Name     string  `gorm:"size:100;" json:"name"`
	BaseUrl  string  `gorm:"size:255;not null" json:"base_url"`
	UserChar string  `gorm:"size:1;default:'@';not null" json:"user_char"`
	Users    []*User `gorm:"many2many:social_media_profiles" json:"users"`
}

type SocialMediaProfile struct {
        ID            uint        `json:"id"`
	UserID        uint        `gorm:"primaryKey" json:"user_id"`
	SocialMediaID uint        `gorm:"primaryKey" json:"social_media_id"`
	User          User        `json:"user"`
	SocialMedia   SocialMedia `json:"social_media"`
	Nickname      string      `gorm:"size:100;not null;unique" json:"nickname"`
}
