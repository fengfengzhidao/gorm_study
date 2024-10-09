// /5.一对一关系.go
package main

import (
	"gorm_study_new/global"
	"gorm_study_new/models"
)

func ManyToMany() {
	// 创建一篇文章，新增tag
	//err := global.DB.Create(&models.ArticleModel{
	//	Title: "文章1",
	//	TagList: []models.TagModel{
	//		{Title: "go"},
	//		{Title: "python"},
	//	},
	//}).Error
	//fmt.Println(err)
	// 创建一篇文章，选择tag
	//var tagIDList = []int{2}
	//var tagList []models.TagModel
	//global.DB.Find(&tagList, "id in ?", tagIDList)
	//
	//err := global.DB.Create(&models.ArticleModel{
	//	Title:   "文章2",
	//	TagList: tagList,
	//}).Error
	//fmt.Println(err)

	// 查文章的时候，把对应的标签带出来
	//var articleList []models.ArticleModel
	//global.DB.Preload("TagList").Find(&articleList)
	//fmt.Println(articleList)

	// 将文章2的标签更新为1,2
	var article models.ArticleModel
	global.DB.Take(&article, 2)

	global.DB.Model(&article).Association("TagList").Replace([]models.TagModel{
		{ID: 1},
		{ID: 2},
	})

}

func main() {
	global.Connect()
	//global.Migrate()
	ManyToMany()
}
