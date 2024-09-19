package handler

import (
	"github.com/MikelSot/trival-proyect/infrastructure/handler/user"
	"github.com/MikelSot/trival-proyect/model"
)

func InitRouter(config model.ConfigRouter) {
	// U
	user.NewRouter(config)
}
