package main

import (
	"Crud-go/src/configuration/database/mongodb"
	"Crud-go/src/configuration/logger"
	"Crud-go/src/controller"
	"Crud-go/src/controller/routes"
	"Crud-go/src/model/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"


)


func main(){


	logger.Info("Esta começando a aplicação")
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	mongodb.InitConnection()
	
	//init dependecies
	service:= service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service) 
		
	router:= gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err:= router.Run(":8080"); err != nil{
		log.Fatal(err)
	}

}