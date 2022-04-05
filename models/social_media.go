package models

type SocialMedia struct {
	ID       uint    `json:"id"`
	Name     string  `gorm:"size:100;not null" json:"name,omitempty"`
	BaseUrl  string  `gorm:"size:255;not null" json:"base_url,omitempty"`
	UserChar string  `gorm:"size:1;default:'@';not null" json:"user_char,omitempty"`
	Users    []*User `gorm:"many2many:social_media_profiles" json:"users,omitempty"`
}

type SocialMediaProfile struct {
    ID            *uint        `json:"id,omitempty"`
	UserID        *uint        `gorm:"primaryKey" json:"user_id,omitempty"`
	SocialMediaID *uint        `gorm:"primaryKey" json:"social_media_id,omitempty"`
	User          *User        `json:"user,omitempty"`
	SocialMedia   SocialMedia `json:"social_media,omitempty"`
	Nickname      string      `gorm:"size:100;not null;unique" json:"nickname,omitempty"`
}
