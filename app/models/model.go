package model

import (
	gorm "github.com/jinzhu/gorm"

	"app/database"
)

type User struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func Migrate() {
  db := database.InitDB()
  db.AutoMigrate(&Product{})
}

func GetUsers() ([]User, error) {
  db := database.InitDB()

	var users []User
	err := db.Find(&users).Error

	return users, err
}

func PostUser(name string) (User, error) {
  db := database.InitDB()

  user := User{Name: name}
  err := db.Create(&user).Error

  return user, err
}