// global/enter.go
package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_study_new/models"
	"log"
)

var DB *gorm.DB

func Migrate() {
	err := DB.AutoMigrate(
		&models.UserModel{},
		&models.UserDetailModel{},
		&models.GirlModel{},
		&models.BoyModel{},
		&models.ArticleModel{}, &models.TagModel{})
	if err != nil {
		log.Fatalf("数据库迁移失败 %s", err)
	}
	fmt.Println("数据库迁移成功")

}

func Connect() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_db_new?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}
