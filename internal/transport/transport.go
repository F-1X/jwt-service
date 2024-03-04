package transport

type HTTP interface {
	AccessHandler(userID string) (accessToken string, refreshToken string, err error)
	RefreshHandler(refreshToken string) (accessToken string, err error)
}



type GRPC interface {
	AccessHandler(userID string) (accessToken string, refreshToken string, err error)
	RefreshHandler(refreshToken string) (accessToken string, err error)
}
