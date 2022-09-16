package database

import (
	"go-payment-system/models"
	"log"

	//"gorm.io/driver/mysql"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	// Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{}) // mysql
	Instance, dbError = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{
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

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
