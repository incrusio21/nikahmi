package auth

type Auth struct {
	Doctype   string `gorm:"primary_key;"`
	Name      string `gorm:"primary_key;"`
	Fieldname string `gorm:"primary_key;"`
	Password  string
}

func (u *Auth) TableName() string {
	return "__auth"
}
