package app

import (
	"jwt-service/internal/config"
	"jwt-service/internal/repository"
	"jwt-service/internal/service"
	"jwt-service/internal/transport/grpc"

	jwtRepository "jwt-service/internal/repository/mongo"
	jwtService "jwt-service/internal/service/jwt"
)

type serviceProdiver struct {
	gRPCServer grpc.ServerHandlers

	jwtRepository repository.JWTRepository
	jwtService    service.JWTService
}

func NewServiceProvider(cfg config.ServiceConfig) *serviceProdiver {

	jwtRepository := jwtRepository.New(cfg.Service.JWT.DB.Mongo)

	jwtService := jwtService.New(jwtRepository, cfg.Service.JWT.Secret)

	gRPCServer := grpc.New(cfg.Service.JWT.GRPC, jwtService)

	return &serviceProdiver{
		gRPCServer:    gRPCServer,
		jwtRepository: jwtRepository,
		jwtService:    jwtService,
	}
}
