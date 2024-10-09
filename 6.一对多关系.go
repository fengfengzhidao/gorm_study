// /5.一对一关系.go
package main

import (
	"gorm_study_new/global"
	"gorm_study_new/models"
)

func oneToMany() {
	// 来了一个女神，自带俩舔狗
	//err := global.DB.Create(&models.GirlModel{
	//	Name: "小露",
	//	BoyList: []models.BoyModel{
	//		{Name: "一号"},
	//		{Name: "二号"},
	//	},
	//}).Error
	//fmt.Println(err)
	// 来了一个女神，选了一个舔狗
	//err := global.DB.Create(&models.GirlModel{
	//	Name: "小张",
	//	BoyList: []models.BoyModel{
	//		{ID: 2},
	//	},
	//}).Error
	//fmt.Println(err)
	// 来了一个舔狗，选了一个女神
	//err := global.DB.Create(&models.BoyModel{
	//	Name:   "三号",
	//	GirlID: 2,
	//}).Error
	//fmt.Println(err)

	// 查询小张的舔狗
	//var girl models.GirlModel
	//global.DB.Preload("BoyList").Take(&girl, "name = ?", "小张")
	//fmt.Println(girl.Name, girl.BoyList, len(girl.BoyList))
	//
	//girl = models.GirlModel{}
	//global.DB.Preload("BoyList", "name = ?", "三号").Take(&girl, "name = ?", "小张")
	//fmt.Println(girl.Name, girl.BoyList, len(girl.BoyList))

	// 专门查关联
	//girl = models.GirlModel{}
	//global.DB.Take(&girl, "name = ?", "小张")
	//var boyList []models.BoyModel

	//global.DB.Model(&girl).Association("BoyList").Find(&boyList)
	//fmt.Println(boyList)

	// 小张的舔狗2号，3号不舔了，换成了1号
	//girl := models.GirlModel{}
	//global.DB.Take(&girl, "name = ?", "小张")
	//global.DB.Model(&girl).Association("BoyList").Replace([]models.BoyModel{
	//	{ID: 1},
	//})

	// 都不舔了
	//girl := models.GirlModel{}
	//global.DB.Take(&girl, "name = ?", "小张")
	//global.DB.Model(&girl).Association("BoyList").Clear()

	// 1，3号又开始舔小张了
	//girl := models.GirlModel{}
	//global.DB.Take(&girl, "name = ?", "小张")
	//global.DB.Model(&girl).Association("BoyList").Append([]models.BoyModel{
	//	{ID: 1},
	//	{ID: 3},
	//})

	// 只有3号退出了
	girl := models.GirlModel{}
	global.DB.Take(&girl, "name = ?", "小张")
	global.DB.Model(&girl).Association("BoyList").Delete([]models.BoyModel{
		{ID: 3},
	})

}

func main() {
	global.Connect()
	global.Migrate()
	oneToMany()
}
