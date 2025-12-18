package main

import (
	"log"
	"Crud-go/src/configuration/logger"
	"Crud-go/src/controller/routes"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"Crud-go/src/model/service"
	"Crud-go/src/controller"

)


func main(){


	logger.Info("Esta começando a aplicação")
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	
	//init dependecies
	service:= service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service) 
		
	router:= gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err:= router.Run(":8080"); err != nil{
		log.Fatal(err)
	}

}