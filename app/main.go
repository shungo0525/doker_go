package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// [ユーザ名]:[パスワード]@tcp([dbのコンテナ名])/[データベース名]
	dbconf := "root:password@tcp(mysql)/development"
	db, err := sql.Open("mysql", dbconf)
	// 接続が終了したらクローズする
	defer db.Close()

	if err != nil {
			fmt.Println(err.Error())
	}

	err = db.Ping()

	if err != nil {
			fmt.Println("データベース接続失敗")
			fmt.Println(err)
			return
	} else {
			fmt.Println("データベース接続成功")
	}
}
