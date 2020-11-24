package Migrate

import (
	Database "api/Database"
	Model "api/Migration"
	"fmt"
)

func MigrationProject() {
	db, err := Database.Open()
	db.AutoMigrate(&Model.User{})
	//1
	/*db.Migrator().AlterColumn(&Model.User{}, "Age")
	db.Migrator().RenameColumn(&Model.User{}, "Age", "Ages")*/
	//2
	/*db.Migrator().DropColumn(&Model.User{}, "Age")
	db.Migrator().RenameColumn(&Model.User{}, "Ages", "Age")*/
	//3
	/*db.Migrator().DropColumn(&Model.User{}, "CreateAt")
	db.Migrator().RenameColumn(&Model.User{}, "CreateAt", "CreatedAt")*/
	//4
	/*db.Migrator().DropColumn(&Model.User{}, "Active")
	db.Migrator().RenameColumn(&Model.User{}, "Status", "Active")*/
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("end")
}
