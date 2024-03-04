package model

type TokenRequest struct {
	UUID string `json:"uuid_id" bson:"user_id,omitempty"`
}

type TokenRequestRefresh struct {
	UUID         string `json:"uuid_id" bson:"user_id,omitempty"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token,omitempty"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token" bson:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token,omitempty"`
}
