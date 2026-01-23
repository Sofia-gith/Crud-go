package converter

import (
	"Go2/src/model/repository/entity"
	"Go2/src/model"


)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.userEntity {
	return &entity.userEntity{
		ID: 	 domain.GetID(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}

	
}