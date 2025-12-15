package controller

import (

	"fmt"
	"Crud-go/src/configuration/rest_err"
	"Crud-go/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("there are some incorrect filds, error=%s", err.Error))

			c.JSON(restErr.Code, restErr)
			return
	}

	fmt.Println(userRequest)
}