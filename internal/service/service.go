package service

import (
	"context"
	"jwt-service/internal/model"
)

type JWTService interface {
	GenerateToken(ctx context.Context, access *model.AccessToken)
	RefreshToken(ctx context.Context, refresh *model.RefreshToken)
}
