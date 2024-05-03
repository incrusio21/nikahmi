package model

import (
	"gorm.io/gorm"
)

type Auth struct {
	Doctype              string `gorm:"primary_key;"`
	Name                 string `gorm:"primary_key;"`
	Fieldname            string `gorm:"primary_key;"`
	Password             string
	SimultaneousSessions int `gorm:"-"`
}

func (u *Auth) TableName() string {
	return "__auth"
}

func (u *Auth) AfterFind(tx *gorm.DB) (err error) {
	// Custom logic after finding a user
	if u.Doctype == "User" {
		user := User{Name: u.Name}
		tx.Model(&User{}).Select("simultaneous_sessions").First(&user)
		u.SimultaneousSessions = user.SimultaneousSessions
	}

	return
}
