// models/user_model.go
package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID              int
	Name            string           `gorm:"size:16;not null;unique"`
	Age             int              `gorm:"default:18"`
	UserDetailModel *UserDetailModel `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL"`
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type UserDetailModel struct {
	ID        int
	UserID    int        `gorm:"unique"` // 一对一的情况下，需要加上唯一约束
	UserModel *UserModel `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL"`
	Email     string     `gorm:"size:32"`
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("创建的钩子函数")
	return nil
}
func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("更新的钩子函数")
	return nil
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("删除的钩子函数")
	return nil
}

func (u *UserModel) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("查询钩子")
	return
}
