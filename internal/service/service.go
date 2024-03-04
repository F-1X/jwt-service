package service

import (
	"context"
	"jwt-service/pkg/jwtGRPC"
)

type JWTService interface {
	GenerateToken(ctx context.Context, UUID *jwtGRPC.TokenRequest) (*jwtGRPC.TokenResponse, error)
	RefreshToken(ctx context.Context, token *jwtGRPC.TokenRequestRefresh) (*jwtGRPC.TokenResponse, error)
}
