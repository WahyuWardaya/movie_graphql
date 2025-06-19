package config

import (
	"fmt"
	"log"
	"movie_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(localhost:3306)/movie_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to MySQL database")

	// AutoMigrate semua tabel (opsional tapi direkomendasikan)
	err = DB.AutoMigrate(
    	&models.Users{},
		&models.Roles{},
		&models.Movies{},
		&models.Broadcast{},
		&models.Statuses{},
		&models.Genres{},
		&models.Actors{},
	)
	if err != nil {
		log.Fatal("Failed to migrate models:", err)
	}
}
