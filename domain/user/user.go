package user

import "github.com/MikelSot/trival-proyect/model"

type DocService interface {
	GetRandomUser() (model.Results, error)
}

type UseCase interface {
	GetUniqueUsers() (model.Users, error)
}
