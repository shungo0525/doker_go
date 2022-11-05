package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id     int
	Name   string
}

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

	// ユーザーデータ取得 & データ出力
	rows := getRows(db)
	printUsers(rows)
}

func getRows(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println("Error")
		panic(err.Error())
	}
	return rows
}

func printUsers(rows *sql.Rows) {
	user := User{}
	var result []User
	for rows.Next() {
		error := rows.Scan(&user.Id, &user.Name)
		if error != nil {
			fmt.Println("scan error")
		} else {
			result = append(result, user)
		}
	}
	fmt.Println(result)
}
