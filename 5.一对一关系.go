// /5.一对一关系.go
package main

import (
	"gorm_study_new/global"
	"gorm_study_new/models"
)

func oneToOne() {
	// 创建用户，连带着创建用户详情
	//err := global.DB.Create(&models.UserModel{
	//	Name: "枫枫",
	//	Age:  25,
	//	UserDetailModel: &models.UserDetailModel{
	//		Email: "枫枫@qq.com",
	//	},
	//}).Error
	//fmt.Println(err)

	// 创建用户详情，关联用户
	//err := global.DB.Create(&models.UserDetailModel{
	//	Email:     "xxx@qq.com",
	//	UserModel: &models.UserModel{ID: 15},
	//	UserID: 15,
	//}).Error
	//fmt.Println(err)

	// 通过用户详情查用户  正查
	//var id = 15
	//var detail models.UserDetailModel
	//global.DB.Preload("UserModel").Take(&detail, "user_id = ?", id)
	//fmt.Println(detail.Email, detail.UserModel.Name)

	//反查
	//var user models.UserModel
	//global.DB.Preload("UserDetailModel").Take(&user, id)
	//fmt.Println(user.Name, user.UserDetailModel.Email)

	// 级联删除
	//var user models.UserModel
	//global.DB.Unscoped().Take(&user, 14)
	//global.DB.Unscoped().Delete(&user)
	//global.DB.Unscoped().Select("UserDetailModel").Delete(&user)

	// set null
	//var user models.UserModel
	//global.DB.Unscoped().Take(&user, 15)
	//global.DB.Model(&user).Association("UserDetailModel").Clear()
	//global.DB.Unscoped().Delete(&user)

	var user models.UserModel
	global.DB.Unscoped().Take(&user, 17)
	global.DB.Unscoped().Delete(&user)
}

func main() {
	global.Connect()
	global.Migrate()
	//oneToOne()
}
