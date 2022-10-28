package database

import (
	"go-payment-system/models"
	"log"

	//"gorm.io/driver/mysql"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	// Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{}) // mysql
	Instance, dbError = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // info 等級的 log
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 讓 gorm 不要再表名後面加上 s

		},
	}) // sqlserver

	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() { // gorm 會自動建立 table
	Instance.AutoMigrate(&models.FourthPartSystemUser{})
	Instance.AutoMigrate(&models.FourthPartyWithdraw{})
	log.Println("Database Migration Completed!")
}
