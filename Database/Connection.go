package Database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1412"
	dbname   = "testgorm"
)

func Open() (*gorm.DB, error) {
	dsn := "user=postgres password=1412 dbname=testgorm port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}
	DB = db
	return DB, err
}
