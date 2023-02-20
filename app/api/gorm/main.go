package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
)

type User struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func main() {
	db := initDB()

	var users []User
	db.Find(&users)

	fmt.Println(users)

	defer db.Close()
}

func initDB() *gorm.DB {
	// 接続情報を設定
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	HOST := "tcp(mysql)"
	DBNAME := "development"

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME

	// DBに接続
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}
