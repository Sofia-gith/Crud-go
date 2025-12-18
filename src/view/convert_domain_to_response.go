package view

import (
	"Crud-go/src/controller/model/response"
	"Crud-go/src/model"
	
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {

	return response.UserResponse{
		ID:        "",
		Name:      userDomain.GetName(),
		Email:     userDomain.GetEmail(),
		Age:	   userDomain.GetAge(),
	
	}

}