package user

import (
	"database/sql"
)

type User struct {
	Name      string `gorm:"primary_key;column:name"`
	Email     string
	FirstName string
	Owner     sql.NullString
}

func (u *User) TableName() string {
	return "tab_user"
}
