package models

import (
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
