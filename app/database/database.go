package database

import (
	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
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

func ConnectCheck() {
  db := InitDB()
  defer db.Close()
}
