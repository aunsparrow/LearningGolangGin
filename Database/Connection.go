package Database

import (
	"fmt"
	"os"

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
	dsn := fmt.Sprintf(`user=%v password=%v dbname=%v port=%v `+
		`sslmode=disable TimeZone=Asia/Bangkok`, os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASENAME"), os.Getenv("PG_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}
	DB = db
	return DB, err
}
