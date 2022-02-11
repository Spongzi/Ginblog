package model

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Cid     int    `json:"cid"`
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"` // 文章描述
	Content string `gorm:"type:longtext" json:"content"`  // 文章内容
	Img     string `gorm:"type:varchar(100)" json:"img"`  // 文章图片
}
