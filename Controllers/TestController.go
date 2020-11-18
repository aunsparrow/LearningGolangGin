package Controllers

import (
	Database "api/Database"
	DbModel "api/Migration"
	Model "api/models/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	var user DbModel.User
	db, err := Database.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	//test commit
	var model Model.UserModelDb
	db.Model(&user).Find(&model)

	var res interface{}
	if model.UserId.String == "" {
		res = Model.UserModel{model.UserId.String, model.FirstName.String, model.LastName.String,
			int(model.Age.Int32), int(model.Status.Int32), model.CreateAt.String}
	}
	res = nil

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcomsdasdsadasdsadsae To API",
		"resualt": res,
	})
	return
}
