package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_api_laundry"))

	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&Pesanan{})

	DB = database

}