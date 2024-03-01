package repository

import (
	"context"

	"github.com/F-1X/jwt-service/internal/model"
)

type JWTRepository interface {
	refreshToken(ctx context.Context, userUUID string, token *model.RefreshToken) error
}
