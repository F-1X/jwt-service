package jwt

import (
	"jwt-service/internal/repository"
	"jwt-service/pkg/jwtGRPC"
)

type JWTService struct {
	repo   repository.JWTRepository
	secret string
	jwtGRPC.UnimplementedJWTServiceServer
}

func New(JWTRepository repository.JWTRepository, secret string) *JWTService {
	
	return &JWTService{

		repo:   JWTRepository,
		secret: secret,
	}
}
