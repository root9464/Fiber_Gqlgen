package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Data struct {
	DB *gorm.DB
}

var Db Data
func ConnectToDB() (Data, error) {
	dsn := "postgres://postgres.fjfcexhrmlodechtmjxu:qwest7q8q1579@aws-0-ap-northeast-1.pooler.supabase.com:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных")
	}
	Db = Data{
		DB: db,
	}
	return Db, nil
}