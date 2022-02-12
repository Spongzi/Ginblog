package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Cid     int    `gorm:"type:int; not null" json:"cid"`
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"` // 文章描述
	Content string `gorm:"type:longtext" json:"content"`  // 文章内容
	Img     string `gorm:"type:varchar(100)" json:"img"`  // 文章图片
}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticle TODO 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	err = db.Debug().Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

// GetArtInfo TODO 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	//fmt.Println("我是art", art)
	err = db.Debug().Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		fmt.Println(err)
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// GetCatArt TODO 查询分类下的所有文章
func GetCatArt(pageSize, pageNum, id int) ([]Article, int) {
	var catArtList []Article
	err = db.Debug().Preload("Category").
		Limit(pageSize).
		Offset((pageNum-1)*pageSize).
		Where("cid = ?", id).
		Find(&catArtList).
		Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return catArtList, errmsg.SUCCESS
}

// EditArticle 编辑分类
func EditArticle(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Debug().Model(&Article{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArticle 删除分类
func DeleteArticle(id int) int {
	err = db.Debug().Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
