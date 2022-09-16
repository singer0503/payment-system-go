package models

import (
	"time"
)

func (FourthPartyWithdraw) TableName() string { // 給 gorm 用的 TableName, 讓開頭變大寫
	return "FourthPartyWithdraw" // 資料表名稱 gorm 會抓取這個名子
}

type FourthPartyWithdraw struct {
	//gorm.Model           // gorm.Model 規範為模型添加了一些默認屬性，例如 id、創建日期、修改日期和刪除日期。
	Id           uint       `gorm:"primarykey;column:Id"` // 主鍵 + 欄位名稱
	DataType     int        `gorm:"column:DataType"`      // 若不設定欄位名稱 gorm 預設會變成 data_type
	OrderId      string     `gorm:"column:OrderId"`
	WorkerId     int        `gorm:"column:WorkerId"`
	PaymentType  string     `gorm:"column:PaymentType"`
	ActualAmount float64    `gorm:"column:ActualAmount"`
	CreatedAt    *time.Time `gorm:"column:CreatedAt"`
}

// // 註冊時，將密碼加密
// func (user *FourthPartSystemUser) HashPassword(password string) error {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(bytes)
// 	return nil
// }

// // 登入時，比對密碼
// func (user *FourthPartSystemUser) CheckPassword(providedPassword string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
