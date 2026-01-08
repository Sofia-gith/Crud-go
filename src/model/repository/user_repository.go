package repository

import (
	"Crud-go/src/configuration/rest_err"
	"Crud-go/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func newUserRepository(
	database *mongo.Database,
)UserRepository{
	return *UserRepository{
		database,
	}
}

type userRepositoy struct{
	databaseConection  *mongo.Database
}

type UserRepository interface{
	CreateUser(
		userDomain model.UserDomainInterface,
	)(model.UserDomainInterface, *rest_err.RestErr)

}