// /2.gorm小试牛刀.go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_db_new"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(db)
	var userList []User
	db.Find(&userList)
	fmt.Println(userList)

}
