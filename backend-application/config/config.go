package config

import (
	"flight-app/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var flightBookingAppDBConnector *gorm.DB

func DatabaseDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
}

func SetDBConnector(db *gorm.DB) {
	flightBookingAppDBConnector = db
}

func GetDBConnector() *gorm.DB {
	return flightBookingAppDBConnector
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(DatabaseDsn()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Airport{},
		&models.Booking{},
		&models.Routes{},
		&models.DetailClass{},
		&models.Flight{},
		&models.Passenger{},
		&models.Payment{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}
