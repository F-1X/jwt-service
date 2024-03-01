package main

import (
	"fmt"

	"jwt-service/pkg/config"
	"jwt-service/pkg/repository"
	"time"
)

type Token struct {
	RefreshToken string    `json:"refresh_token"`
	AccessToken  string    `json:"access_token"`
	UserID       string    `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func main() {

	cfg := config.ReadConfig()

	fmt.Println(cfg.Service.JWT.DB.Mongo)

	repository := repository.InitRepository(&cfg)
	fmt.Println("rep", repository)

}
