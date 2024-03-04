package jwt

import (
	"context"
	"errors"
	"log"
	"time"

	"jwt-service/internal/model"
	"jwt-service/pkg/jwtGRPC"


	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *JWTService) RefreshToken(ctx context.Context, token *jwtGRPC.TokenRequestRefresh) (*jwtGRPC.TokenResponse, error) {

	log.Println("new refresh token:", token)
	filter := bson.M{"uuid": token.Uuid, "refresh_token": token.RefreshToken}

	refreshToken, ok := filter["refresh_token"].(string)
	if !ok {
		log.Println("Refresh token is not a string")
		return &jwtGRPC.TokenResponse{}, nil
	}

	TokenRequest := &model.TokenRequestRefresh{
		UUID:         token.Uuid,
		RefreshToken: token.RefreshToken,
	}

	if err := verifyToken(refreshToken); err != nil {
		log.Println("Failed to verify refresh token:", err)
		return &jwtGRPC.TokenResponse{}, nil
	}

	ok, err := s.repo.RefreshToken( TokenRequest.UUID, refreshToken)
	if err != nil {
		log.Println(err)

	}
	if ok {
		log.Println("exist ok")
	}

	return &jwtGRPC.TokenResponse{}, nil
}
func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte{}, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token is invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if expirationTime.Before(time.Now()) {
			return errors.New("token has expired")
		}
	}

	return nil
}
