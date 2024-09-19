package bootstrap

import (
	"log"

	"github.com/MikelSot/trival-proyect/infrastructure/handler"
	"github.com/MikelSot/trival-proyect/model"
)

func Run() {
	app := newHttp()

	handler.InitRouter(model.ConfigRouter{
		App:           app,
		RandomUserUrl: "https://randomuser.me/api/?results=5000",
	})

	log.Println("Starting server")
	app.Logger.Fatal(app.Start(":5000"))
}
