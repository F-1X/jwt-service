package model

type AccessToken struct {
	TokenString string `json:"token_string"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	Expiration  int64  `json:"expiration"`
}

type RefreshToken struct {
	TokenString string `json:"token_string"`
	UserID      string `json:"user_id"`
	Expiration  int64  `json:"expiration"`
}
