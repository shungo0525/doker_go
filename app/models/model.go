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

var db = database.InitDB()

func Migrate() {
  db.AutoMigrate(&Product{})
}

func GetUsers() ([]User, error) {
	var users []User
	err := db.Find(&users).Error

	return users, err
}

func PostUser(name string) (User, error) {
  user := User{Name: name}
  err := db.Create(&user).Error

  return user, err
}

func PutUser(id int, name string) (User, error) {
  err := db.Model(&User{}).Where("id = ?", id).Update("name", name).Error

  user := User{}
  db.Where("id = ?", id).Take(&user)

  return user, err
}