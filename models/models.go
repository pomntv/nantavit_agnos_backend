package models

import "gorm.io/gorm"

type Password_recommendation struct {
	gorm.Model
	Init_password string `json:"init_password" gorm:"text;not null;default:null`
	Answer        string `json:"answer" gorm:"text;not null;default:null`
}
