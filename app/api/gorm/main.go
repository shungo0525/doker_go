package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"

	"example/api/gorm/models"
)

func main() {
	db := initDB()
	defer db.Close()

	var users []model.User
	db.Find(&users)

	fmt.Println(users)
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
