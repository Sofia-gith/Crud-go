package repository

import(
	"os"
	
	"Crud-go/src/model"
	"Crud-go/src/configuration/rest_err"
	"Crud-go/src/configuration/logger"
)

const(
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

 func (ur *userRepository) CreateUser(
		userDomain model.UserDomainInterface,
	)(model.UserDomainInterface, *rest_err.RestErr){
		logger.Info("Init vreateUser repository ")

		collection_name := os.Getenv(MONGODB_USER_COLLECTION)

		collection := ur.databaseConnection.Collection(collection_name)
	}