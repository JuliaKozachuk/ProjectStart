package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB(postgres_data string) {
	db, err := gorm.Open("postgres", postgres_data)
	//db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=musicstore password=qwerty sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Track{})

	//return db
	DB = db
}
