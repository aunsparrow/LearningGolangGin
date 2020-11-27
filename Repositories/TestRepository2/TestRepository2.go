package TestRepository2

import (
	Database "api/Database"
	Model "api/Migration"
	"fmt"
)

type TestRepository interface {
	GetAll() ([]Model.User, error)
}

var model Model.User

func GetAll() ([]Model.User, error) {
	db, err := Database.Open()
	var res []Model.User = make([]Model.User, 0)
	if err != nil {
		fmt.Println(err)
	}

	db.Model(&model).Find(&res)
	return res, err

}
