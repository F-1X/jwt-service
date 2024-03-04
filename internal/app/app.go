package app

import (
	"jwt-service/internal/config"
)

type App struct {
	cfg             config.ServiceConfig
	serviceProdiver *serviceProdiver
}

func Run() {
	app := &App{}

	app.cfg = config.ReadConfig()

	app.serviceProdiver = NewServiceProvider(app.cfg)

	app.serviceProdiver.gRPCServer.Start()

}
