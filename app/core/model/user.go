package model

import (
	"database/sql"
)

type User struct {
	Name                 string `gorm:"primary_key;column:name"`
	Email                string
	FirstName            string
	Owner                sql.NullString
	SimultaneousSessions int  `gorm:"column:simultaneous_sessions"`
	Auth                 Auth `gorm:"polymorphicType:Doctype;polymorphicId:Name;polymorphicValue:User"`
}

func (u *User) TableName() string {
	return "tab_user"
}
