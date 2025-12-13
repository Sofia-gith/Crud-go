package routes

import (
	controller "Crud-go/src/Controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r*gin.RouterGroup){

	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser/", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.deleteUser)

}