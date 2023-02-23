package model

import (
	gorm "github.com/jinzhu/gorm"
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