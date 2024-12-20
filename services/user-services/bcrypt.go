package userservices

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
)

func HashPassword(password string) (string, error) {
	// Load .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	// Get salt round from environment variable
	saltRound := os.Getenv("saltRound")
	if saltRound == "" {
		return "", fmt.Errorf("saltRound not set in .env file")
	}

	// Convert salt round to integer
	s, err := strconv.Atoi(saltRound)
	if err != nil {
		return "", fmt.Errorf("invalid saltRound value: %v", err)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	// Return hashed password as a byte slice
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
