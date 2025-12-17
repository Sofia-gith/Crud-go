package service


import (
	"Crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
	"Crud-go/src/configuration/logger"
	"fmt"
	"Crud-go/src/model"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr{

	logger.Info("Iniciando createUse Model", zap.String("jorney", "CreateModel"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
		return nil
}