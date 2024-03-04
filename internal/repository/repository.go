package repository

type JWTRepository interface {
	GenerateToken(token string) error
	RefreshToken(userUUID string, token string) (bool, error)
}
