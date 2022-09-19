package models

import (
	"golang.org/x/crypto/bcrypt"
)

func (FourthPartSystemUser) TableName() string { // 給 gorm 用的 TableName, 讓開頭變大寫
	return "FourthPartSystemUser" // 資料表名稱 gorm 會抓取這個名子
}

type FourthPartSystemUser struct {
	//gorm.Model        // gorm.Model 規範為模型添加了一些默認屬性，例如 id、創建日期、修改日期和刪除日期。
	Id       uint   `gorm:"primarykey;column:Id"` // 主鍵 + 欄位名稱
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// 註冊時，將密碼加密
func (user *FourthPartSystemUser) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// 登入時，比對密碼
func (user *FourthPartSystemUser) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
