package models

import (
	"errors"
	"fmt"

	"github.com/skstef/Go-REST-API/db"
	"github.com/skstef/Go-REST-API/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(u.Email, hashPassword)

	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	fmt.Println(passwordIsValid)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
