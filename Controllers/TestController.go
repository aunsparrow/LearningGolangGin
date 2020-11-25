package Controllers

import (
	Model "api/Models/Models"
	Response "api/Models/ResponseModels"
	Service "api/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {

	chGetAll := make(chan []Model.UserModel, 1)
	go func() {
		chGetAll <- Service.TestService()
	}()
	getAll := <-chGetAll
	response := Response.ResponseModel{"true", http.StatusOK, getAll, "Welcome To API"}
	c.JSON(http.StatusOK, response)
	return
}
