package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (User) TableName() string { // 給 gorm 用的 TableName, 讓開頭變大寫
	return "User"
}

type User struct {
	gorm.Model        // gorm.Model 規範為模型添加了一些默認屬性，例如 id、創建日期、修改日期和刪除日期。
	Name       string `json:"name"`
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}

type Role struct {
	Admin string `json:"admin"`
	User  string `json:"user"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
