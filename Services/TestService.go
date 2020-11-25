package Services

import (
	Database "api/Database"
	DbModel "api/Migration"
	Model "api/Models/Models"
	"fmt"
)

func TestService() []Model.UserModel {
	var user DbModel.User
	var res []Model.UserModel = make([]Model.UserModel, 0)
	db, err := Database.Open()
	if err != nil {
		fmt.Println(err.Error())
	}

	db.Model(&user).Find(&res)

	return res
}
