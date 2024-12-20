package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model `json:"gorm.Model"`
	Username   string `gorm:"column:username;not null" form:"username" json:"username" binding:"required,max=100"`
	Password   string `gorm:"column:password;not null" form:"password" json:"password" binding:"required,max=100"`
	Name       string `gorm:"column:name;not null" form:"name" json:"name" binding:"required,max=100"`
	Role       string `gorm:"column:role;default:'user';not null" form:"role" json:"role" binding:"required,max=100"`
	Email      string `gorm:"column:email;type:varchar(100);unique;not null" form:"email" json:"email" binding:"required,email,max=100"`
	Address    string `gorm:"column:address;size:100;not null" form:"address" json:"address" binding:"required,max=100"`
}
