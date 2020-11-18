package Migration

import (
	Database "api/Database"
	"fmt"
)

func MigrationProject() {
	db, err := Database.Open()
	db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("end")
}
