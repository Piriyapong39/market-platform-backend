package user

import (
	"fmt"

	//import services
	userservices "github.com/piriyapong39/market-platform/services/user-services"

	//import database
	db "github.com/piriyapong39/market-platform/config"
)

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Is_seller bool   `json:"is_seller"`
}

func _userRegister(user User) (string, error) {

	//connect to database
	db, err := db.Connection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	//find exists email
	var exists bool
	if err = db.QueryRow(`
        SELECT EXISTS(SELECT 1 FROM tb_users WHERE email = $1)
    `, user.Email).Scan(&exists); err != nil {
		return "", err
	}
	if exists {
		return "", fmt.Errorf("email already exists")
	}

	//hash password
	hashedPassword, err := userservices.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	//insert user data in database
	if _, err = db.Exec(`
		INSERT INTO tb_users (email, password, first_name, last_name)
		VALUES ($1, $2, $3, $4)
	`, user.Email, hashedPassword, user.FirstName, user.LastName); err != nil {
		return "", err
	}

	return "register successfully", nil
}

func _userLogin(user User) (string, error) {
	// declare variable
	userData := new(User)

	//connect database
	db, err := db.Connection()
	if err != nil {
		return "", err
	}

	if err := db.QueryRow(
		`
			SELECT 
				email,
				password,
				first_name,
				last_name,
				is_seller
			FROM 
				tb_users
			WHERE 1=1
				AND email=$1
		`, user.Email).
		Scan(&userData.Email, &userData.Password, &userData.FirstName,
			&userData.LastName, &userData.Is_seller); err != nil {
		return "", err
	}

	isMatch := userservices.CheckPasswordHash(user.Password, userData.Password)
	if !isMatch {
		return "", fmt.Errorf("wrong password please try again")
	}

	token, err := userservices.GenerateToken(userData.Email, userData.FirstName,
		userData.LastName, userData.Is_seller)
	if err != nil {
		return "", err
	}
	return token, nil
}
