package models

import "time"

type Post struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `gorm:"size:255;not null;unique" json:"title"`
	Content    string    `gorm:"size:512;not null;" json:"content"`
	CoverImage string    `gorm:"size:512" json:"cover_img"`
	Author     User      `json:"author"`
	AuthorID   uint32    `gorm:"not null" json:"author_id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
