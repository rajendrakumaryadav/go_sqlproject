package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey" json:"id" xml:"id"`
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}
