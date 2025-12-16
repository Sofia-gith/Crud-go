package model

import (
	"Crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
	"Crud-go/src/configuration/logger"
	"fmt"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr{

	logger.Info("Iniciando createUse Model", zap.String("jorney", "CreateModel"))

	ud.EncryptPassword()

	fmt.Println(ud)
		return nil
}