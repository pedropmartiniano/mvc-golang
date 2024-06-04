package controller

import (

	"github.com/gin-gonic/gin"
	"github.com/pedropmartiniano/mvc-golang/src/configuration/validation"
	"github.com/pedropmartiniano/mvc-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		httpErr := validation.ValidateUserError(err)
			c.JSON(httpErr.Code, httpErr)
		return
	}
}
