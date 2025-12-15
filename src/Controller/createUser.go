package controller

import (
	
	"log"
	"fmt"
	"Crud-go/src/controller/model/response"
	"Crud-go/src/controller/model/request"
	"Crud-go/src/configuration/validation"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error tryng to marrshall object, error: %s", err.Error())
		restErr := validation.ValidatorError(err)

			c.JSON(restErr.Code, restErr)
			return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID: 		"test",
		Name: 	userRequest.Name,
		Email: 	userRequest.Email,
	}
	c.JSON(201, response)
}