package Controllers

import (
	Model "api/Migration"
	Response "api/Models/ResponseModels"
	Service "api/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {

	chGetAll := make(chan []Model.User, 1)
	go func() {
		a, _ := Service.TestService()
		chGetAll <- a
	}()
	getAll := <-chGetAll
	response := Response.ResponseModel{"true", http.StatusOK, getAll, "Welcome To API"}
	c.JSON(http.StatusOK, response)
	return
}
