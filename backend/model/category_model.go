package model
// import (
// 	"gorm.io/gorm"
// )

type Category struct {
 ID          int    `gorm:"column:id_category; primary_key; not null" json:"id"`
 Name   string `gorm:"column:name" json:"category"`
}