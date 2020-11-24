package Migration

import (
	Database "api/Database"
	"fmt"
)

func MigrationProject() {
	db, err := Database.Open()
	db.AutoMigrate(&User{})
	db.Migrator().DropColumn(&User{}, "Ages")
	db.Migrator().RenameColumn(&User{}, "Age", "Ages")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("end")
}
