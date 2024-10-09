// /1.原生sql.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_db_new"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("dsn格式错误 %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败 %s", err)
	}
	fmt.Println(db)

	//res, err := db.Exec("insert into users (name, age, email) values ('张三', 18, 'zhangsan@example.com'), ('李四', 20, 'lisi@example.com')")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res)

	rows, _ := db.Query("select id, name from users")
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Println(id, name, err)
	}
	var id int
	var name string
	err = db.QueryRow("select id, name from users").Scan(&id, &name)
	fmt.Println(id, name, err)
}
