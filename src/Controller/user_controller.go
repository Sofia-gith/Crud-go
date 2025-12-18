package controller

import (
	"Crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

func newUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface{
	return &userControllerInterface{}
}

type UserControllerInterface interface{
	FindUserById(c *gin.Context)
	FindByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

}

type userControllerInterface struct{
	service service.UserDomainService
}