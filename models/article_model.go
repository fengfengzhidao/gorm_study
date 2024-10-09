// models/article_model.go
package models

type ArticleModel struct {
	ID      int
	Title   string     `gorm:"size:32"`
	TagList []TagModel `gorm:"many2many:article_tags;"`
}

type TagModel struct {
	ID          int
	Title       string         `gorm:"size:32"`
	ArticleList []ArticleModel `gorm:"many2many:article_tags;"`
}
