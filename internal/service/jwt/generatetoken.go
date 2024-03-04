package jwt

import (
	"context"
	"log"
	"time"

	"jwt-service/pkg/jwtGRPC"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	uuid string
	jwt.StandardClaims
}

func (s *JWTService) GenerateToken(ctx context.Context, req *jwtGRPC.TokenRequest) (*jwtGRPC.TokenResponse, error) {
	log.Println("ner genratin request", req)
	accessToken, err := generateToken(req.Uuid, time.Duration(15*time.Minute)).SignedString([]byte(s.secret))
	if err != nil {
		log.Println(err)
		return &jwtGRPC.TokenResponse{AccessToken: "", RefreshToken: ""}, nil
	}

	if err := s.repo.GenerateToken(accessToken); err != nil {
		log.Println(err)
	}

	refreshToken, err := generateToken(req.Uuid, time.Duration(1*time.Hour)).SignedString([]byte(s.secret))
	if err != nil {
		log.Println(err)
		return &jwtGRPC.TokenResponse{AccessToken: "", RefreshToken: ""}, nil
	}
	log.Println("fin gen",accessToken,refreshToken)
	return &jwtGRPC.TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func generateToken(uuid string, ttl time.Duration) *jwt.Token {
	claims := &JWTClaims{
		uuid: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token
}
