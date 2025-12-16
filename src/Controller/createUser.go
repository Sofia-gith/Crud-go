package controller

import (
	"Crud-go/src/configuration/logger"
	"Crud-go/src/configuration/validation"
	"Crud-go/src/controller/model/request"
	"Crud-go/src/controller/model/response"
	"Crud-go/src/model/user_domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

var(
	userDomainInterface model.UserDomainInterface
)
func CreateUser(c *gin.Context){
	logger.Info("Iniciando createUserController",
	zapcore.Field{
		Key: "journey",
		String: "createUserController",
	})

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error tryng to validate user info", err,
			zapcore.Field{
				Key: "journey",
				String: "createUserController",
			})
		restErr := validation.ValidatorError(err)

			c.JSON(restErr.Code, restErr)
			return
	}

	logger.Info("User created successfully",
			zapcore.Field{
				Key: "journey",
				String: "createUser",
			})
	response := response.UserResponse{
		ID: 		"test",
		Name: 	userRequest.Name,
		Email: 	userRequest.Email,
		Age:  	userRequest.Age,
	}
	c.JSON(http.StatusOK, response)
}