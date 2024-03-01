package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func generateAccessToken(userID string) (string, error) {

	claims := &model.JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(Key)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
