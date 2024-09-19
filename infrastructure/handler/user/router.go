package user

import (
	"github.com/MikelSot/trival-proyect/domain/user"
	"github.com/MikelSot/trival-proyect/infrastructure/service"
	"github.com/MikelSot/trival-proyect/model"
)

const (
	_publicRouter = "/users"
)

func NewRouter(config model.ConfigRouter) {
	h := buildHandler(config)

	publicRoute(config, h)
}

func buildHandler(config model.ConfigRouter) handler {
	useCase := user.New(service.New(config))

	handler := newHandler(useCase)

	return handler
}

func publicRoute(config model.ConfigRouter, h handler) {
	config.App.GET(_publicRouter, h.GetUniqueUsers)
}
