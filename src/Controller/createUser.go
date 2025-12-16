package controller

import (
	"Crud-go/src/configuration/logger"
	"Crud-go/src/configuration/validation"
	"Crud-go/src/controller/model/request"
	"Crud-go/src/model"
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

	
			domain := model.NewUserDomain(
				userRequest.Email,
				userRequest.Password,
				userRequest.Name,
				userRequest.Age,	
			)
			if err := domain.CreateUser(); err != nil{
				c.JSON(http.StatusOK, err)
				return
			}

	logger.Info("User criado com sucesso",
			zapcore.Field{
				Key: "journey",
				String: "createUser",
			})


	c.String(http.StatusOK, "")
}