package userservices

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/joho/godotenv"
)

type UserData struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Is_seller bool   `json:"is_seller"`
}

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

func VerifyToken(tokenString string) (UserData, error) {
	if err := godotenv.Load("./config/.env"); err != nil {
		return UserData{}, err
	}
	JWT_SECRET_KEY := []byte(os.Getenv("JWT_SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET_KEY, nil
	})
	if err != nil || !token.Valid {
		return UserData{}, fmt.Errorf("invalid token: %v", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return UserData{}, fmt.Errorf("invalid claims")
	}
	userData := UserData{
		Email:     claims["email"].(string),
		FirstName: claims["firstName"].(string),
		LastName:  claims["lastName"].(string),
		Is_seller: claims["isSeller"].(bool),
	}
	return userData, nil
}
