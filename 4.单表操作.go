// /4.单表操作.go
package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_study_new/global"
	"gorm_study_new/models"
	"time"
)

func insert() {
	//err := global.DB.Create(&models.UserModel{
	//	Name: "枫枫",
	//	Age:  25,
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 回填式创建
	//user := models.UserModel{
	//	Name: "张三",
	//	Age:  25,
	//}
	//err := global.DB.Create(&user).Error
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(user.ID, user.CreatedAt, user.Name)

	//var userList = []models.UserModel{
	//	{Name: "王五"},
	//	{Name: "李四"},
	//}
	//err := global.DB.Create(&userList).Error
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	err := global.DB.Create(&models.UserModel{
		Name: "枫枫1",
		Age:  25,
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}

func query() {
	var userList []models.UserModel
	global.DB.Find(&userList, "age >= ?", 25)
	fmt.Println(userList)

	//var user models.UserModel
	//global.DB.Take(&user)
	// 根据主键去查
	//global.DB.Take(&user, 4)
	// 根据主键排序查第一个
	//global.DB.Debug().First(&user, 4)
	//fmt.Println(user)
	//global.DB.Debug().Last(&user, 5)
	//fmt.Println(user)

	//user.ID = 6
	//global.DB.Take(&user)
	//fmt.Println(user)

	//var user models.UserModel
	//err := global.DB.Limit(1).Find(&user, 111).Error
	//if err == gorm.ErrRecordNotFound {
	//	fmt.Println("记录不存在")
	//	return
	//}
	//fmt.Println(user.ID)

}

func update() {
	var user = models.UserModel{ID: 3}
	global.DB.Model(&user).Update("age", 0)
	global.DB.Model(&user).UpdateColumn("age", 0)
	fmt.Println(user)
}

func updates() {
	var user = models.UserModel{ID: 3}
	//global.DB.Model(&user).Updates(models.UserModel{
	//	Name: "zhangsan",
	//	Age:  0,
	//})
	//global.DB.Model(&user).Updates(map[string]any{
	//	"age": 0,
	//})
	//fmt.Println(user)
	global.DB.Model(&user).Update("age", gorm.Expr("age + ?", 1))
	fmt.Println(user)
}

func updateSave() {
	var user models.UserModel
	user.ID = 3
	user.Age = 0
	user.Name = "张四"
	user.CreatedAt = time.Now()
	global.DB.Save(&user)
	fmt.Println(user)
}

func Delete() {
	//var user = models.UserModel{ID: 9}
	//global.DB.Delete(&user)

	//global.DB.Delete(&models.UserModel{}, 3)

	//global.DB.Delete(&models.UserModel{}, "name = ?", "枫枫1")

	//global.DB.Delete(&models.UserModel{}, []int{4, 5})

	var user models.UserModel
	global.DB.Unscoped().Take(&user, 3)
	fmt.Println(user)
	global.DB.Unscoped().Delete(&user)
}

func HighQuery() {
	//var user models.UserModel
	//global.DB.Debug().Where("age > ?", 18).Take(&user)
	//fmt.Println(user)

	//var user models.UserModel
	//global.DB.Debug().Where(models.UserModel{
	//	Name: "",
	//	Age:  0,
	//}).Take(&user)
	//fmt.Println(user)

	//var user models.UserModel
	//global.DB.Debug().Where(map[string]any{
	//	"age": 0,
	//}).Take(&user)
	//fmt.Println(user)

	//query := global.DB.Where("age > 18 and name = ?", "枫枫")
	//var user models.UserModel
	//global.DB.Debug().Where(query).Take(&user)
	//fmt.Println(user)

	//var userList []models.UserModel
	//global.DB.Debug().Or("age = 24").Or("name = ?", "枫枫").Find(&userList)
	//fmt.Println(userList)

	//var userList []models.UserModel
	//global.DB.Debug().Not("age > 18").Find(&userList)
	//fmt.Println(userList)

	//var userList []models.UserModel
	//global.DB.Order("age desc").Find(&userList)
	//global.DB.Order("age asc").Order("id desc").Find(&userList)
	//fmt.Println(userList)

	//var nameList []string
	//global.DB.Model(models.UserModel{}).Select("name").Scan(&nameList)
	//fmt.Println(nameList)

	//global.DB.Model(models.UserModel{}).Pluck("name", &nameList)
	//fmt.Println(nameList)

	//type User struct {
	//	Name string
	//	Age int
	//}
	//type User struct {
	//	Label string `gorm:"column:name"`
	//	Value int    `gorm:"column:age"`
	//}
	//var userList []User
	//global.DB.Model(models.UserModel{}).Scan(&userList)
	//
	//fmt.Println(userList)

	//type UserCount struct {
	//	Age   int
	//	Count int
	//}
	//var lis []UserCount
	//global.DB.Model(models.UserModel{}).Group("age").Select("age", "count(id) as count").Scan(&lis)
	//fmt.Println(lis)

	//var ageList []int
	//global.DB.Model(models.UserModel{}).Distinct("age").Pluck("age", &ageList)
	//fmt.Println(ageList)

	//var userList []models.UserModel
	//limit := 2
	//page := 3
	//global.DB.Limit(2).Offset(0).Find(&userList)
	//fmt.Println(userList)
	//global.DB.Limit(2).Offset(2).Find(&userList)
	//fmt.Println(userList)
	//global.DB.Limit(2).Offset(4).Find(&userList)
	//fmt.Println(userList)
	//global.DB.Limit(limit).Offset((page - 1) * limit).Find(&userList)
	//fmt.Println(userList)

	//var userList []models.UserModel
	//global.DB.Scopes(Age18).Find(&userList)
	//global.DB.Scopes(NameIn("枫枫", "王五")).Find(&userList)
	//fmt.Println(userList)

	type User struct {
		Name string
		Age  int
	}
	var users []User
	global.DB.Raw("select name, age from user_models where age > 18").Scan(&users)
	fmt.Println(users)

	global.DB.Exec("update user_models set age = 19 where id = 13")

}
func NameIn(nameList ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("name in ?", nameList)
	}
}

func Age18(tx *gorm.DB) *gorm.DB {
	return tx.Where("age > 18")
}

func main() {

	global.Connect()
	global.Migrate()
	//query()
	//updateSave()
	//update()
	//updates()
	//Delete()
	HighQuery()
}
