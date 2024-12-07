package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Name         string `binding:"required" json:"Name" gorm:"not null;varchar(255)"`
	Selector string `binding:"required" json:"Selector" gorm:"not null;"`
	Url           string `binding:"required" json:"url" gorm:"not null;`
}

type ArticleDTO struct {
	URLs          []string `binding:"required" json:"urls" `
	Selector string   `binding:"required" json:"Selector"`
}
