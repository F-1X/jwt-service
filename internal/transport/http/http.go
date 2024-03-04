package http

import (
	"jwt-service/internal/config"

	"github.com/labstack/echo"
)

// type Token struct {

// 	RefreshToken string    `json:"refresh_token"`
// 	AccessToken  string    `json:"access_token"`
// 	UserID       string    `json:"user_id"`
// 	ExpiresAt    time.Time `json:"expires_at"`
// }

// type HTTP interface {
// 	AccessHandler(c echo.Context, userID string) (err error)
// 	RefreshHandler(refreshToken string) (accessToken string, err error)
// }

// var _ HTTP = (*httpServer)(nil)

type ServerHandlers interface {
	AccessHandler(c echo.Context) error
	RefreshHandler(c echo.Context) (string, error)
}

type Server struct {
	cfg config.HTTP
}

func New(cfg config.HTTP) ServerHandlers {
	return &Server{
		cfg: cfg,
	}
}

// func (s *httpServer) Start() {
// 	e := echo.New()

// 	e.POST("/tokens", s.AccessHandler)
// 	e.POST("/tokens/refresh", s.RefreshHandler)

// 	e.Logger.Fatal(e.Start(":8080"))
// }

func (s *Server) AccessHandler(c echo.Context) error {
	return nil
}

func (s *Server) RefreshHandler(c echo.Context) (string, error) {
	return "", nil
}

// func hashRefreshToken(refreshToken string) (string, error) {
// 	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedToken), nil
// }
