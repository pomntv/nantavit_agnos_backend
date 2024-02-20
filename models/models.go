package models

import "gorm.io/gorm"

type Password_recommendation struct {
	gorm.Model
	Init_password string `json:"init_password" gorm:"text;not null;default:null`
}
type Log struct {
	gorm.Model
	Message string
}
