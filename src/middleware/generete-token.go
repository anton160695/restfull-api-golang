package middleware


import (
	"crud-golang/crud/src/database"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken(userId int, username, name string, duration time.Duration) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"name":     name,
		"exp":      time.Now().Add(duration).Unix(),
	}

	JwtScreet := os.Getenv("SCREET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtScreet))
}

func GenerateToken(user database.Users) (string, error) {
	accessToken, err := createToken(user.Id, user.Username, user.Name, 15*time.Minute)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

