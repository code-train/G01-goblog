package user

import (
	"goblog/app/models"
	"goblog/pkg/password"
)

// User 模型
type User struct {
	models.BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255);default null;unique" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255)" valid:"password"`

	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}


func (user User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}

func (user User) Link() string {
	return ""
}