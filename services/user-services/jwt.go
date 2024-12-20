package userservices

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/joho/godotenv"
)

func GenerateToken(email string, firstName string, lastName string, isSeller bool) (string, error) {

	//import .env file
	if err := godotenv.Load("./config/.env"); err != nil {
		return "", err
	}
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	// Create the Claims
	claims := jwt.MapClaims{
		"email":     email,
		"firstName": firstName,
		"lastName":  lastName,
		"isSeller":  isSeller,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return t, nil

}
