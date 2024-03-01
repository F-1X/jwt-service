package jwt

// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"golang.org/x/crypto/bcrypt"
// )




// var secretKey = []byte("secret-key")

// func createToken(username string) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
// 		jwt.MapClaims{
// 			"username": username,
// 			"exp":      time.Now().Add(time.Hour * 24).Unix(),
// 		})

// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func RefreshHandler(c echo.Context) error {

// 	refreshReq := struct {
// 		RefreshToken string `json:"refresh_token"`
// 	}{}

// 	if err := c.Bind(&refreshReq); err != nil {
// 		return c.JSON(http.StatusBadRequest, "Invalid request")
// 	}

// 	var storedToken Token
// 	collection := client.Database(dbName).Collection(collection)
// 	err := collection.FindOne(context.Background(), bson.M{"refresh_token": refreshReq.RefreshToken}).Decode(&storedToken)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, "Invalid refresh token")
// 	}

// 	// Check if the refresh token is expired
// 	if time.Now().After(storedToken.ExpiresAt) {
// 		return c.JSON(http.StatusUnauthorized, "Refresh token expired")
// 	}

// 	// Generate a new access token
// 	accessToken, err := generateAccessToken(storedToken.UserID)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, "Failed to generate access token")
// 	}

// 	// Return the new access token
// 	return c.JSON(http.StatusOK, map[string]string{"access_token": accessToken})
// }



// func AccessHandler(c echo.Context) error {
// 	return c.JSON(http.StatusOK, "Access token generated successfully")
// }





// func hashRefreshToken(refreshToken string) (string, error) {
// 	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedToken), nil
// }
