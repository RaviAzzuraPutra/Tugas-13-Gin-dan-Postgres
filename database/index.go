package database

import (
	"fmt"
	"formative-13/config/db_config"
	"formative-13/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var errorConnect error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8", db_config.DB_Config().DB_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME, db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE)

	DB, errorConnect = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errorConnect != nil {
		panic("An error occurred while connecting to the database! : " + errorConnect.Error())
	}

	if DB == nil {
		panic("TERJADI KESALAHAN DI DATABASE!!")
	}

	errMigrate := DB.AutoMigrate(&model.Bioskop{})

	if errMigrate != nil {
		panic("An error occurred during database migration : " + errMigrate.Error())
	}

	log.Println("Berhasil Konek Ke Database")

}
