package userservices

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/joho/godotenv"
)

type UserData struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Is_seller bool   `json:"is_seller"`
}

func GenerateToken(id int, email string, firstName string, lastName string, isSeller bool) (string, error) {

	//import .env file
	if err := godotenv.Load("./config/.env"); err != nil {
		return "", err
	}
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	// Create the Claims
	claims := jwt.MapClaims{
		"id":        id,
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

	tokenPart := strings.Split(tokenString, " ")
	if len(tokenPart) != 2 || tokenPart[0] != "Bearer" {
		return UserData{}, fmt.Errorf("invalid token")
	}

	token, err := jwt.Parse(tokenPart[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return UserData{}, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return UserData{}, fmt.Errorf("invalid claims")
	}

	id, ok1 := claims["id"].(float64)
	email, ok2 := claims["email"].(string)
	firstName, ok3 := claims["firstName"].(string)
	lastName, ok4 := claims["lastName"].(string)
	isSeller, ok5 := claims["isSeller"].(bool)

	if !(ok1 && ok2 && ok3 && ok4 && ok5) {
		return UserData{}, fmt.Errorf("invalid claim values")
	}
	// fmt.Println(isSeller)
	return UserData{
		Id:        int(id),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Is_seller: isSeller,
	}, nil
}
