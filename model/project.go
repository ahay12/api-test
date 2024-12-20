package model

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	gorm.Model
	Image       string    `gorm:"column:image;size:100;null" form:"image" json:"image,omitempty"`
	Title       string    `gorm:"column:title;size:100;not null" form:"title" json:"title,omitempty"`
	Description string    `gorm:"column:description;size:1000;not null" form:"description" json:"description,omitempty"`
	Goals       uint      `gorm:"column:goals" form:"goals" json:"goals,omitempty"`
	Fund        uint      `gorm:"column:fund" form:"fund" json:"fund,omitempty"`
	Category    string    `gorm:"column:category;size:50;not null" form:"category" json:"category,omitempty"`
	Tag         string    `gorm:"column:tag;size:100;null" form:"tag" json:"tag,omitempty"`
	View        uint      `gorm:"column:view" form:"view" json:"view,omitempty"`
	Expired     time.Time `gorm:"column:expired" form:"expired" json:"expired,omitempty"`
}
