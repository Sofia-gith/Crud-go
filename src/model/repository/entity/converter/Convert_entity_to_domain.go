package converter

import (

	"Crud-go/src/model"
	"Crud-go/src/model/repository/entity"
)

func ConvertEnityToDomain(
	entity entity.UserEntity,
	) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID)
	return domain
}
