package user

import (
	"database/sql"

	"github.com/incrusio21/nikahmi/app/core/page/auth"
)

type User struct {
	Name      string `gorm:"primary_key;column:name"`
	Email     string
	FirstName string
	Owner     sql.NullString
	Auth      auth.Auth `gorm:"polymorphicType:Doctype;polymorphicId:Name;polymorphicValue:User"`
}

func (u *User) TableName() string {
	return "tab_user"
}
