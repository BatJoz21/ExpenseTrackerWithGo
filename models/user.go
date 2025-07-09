package models

import (
	"errors"

	"example.com/expense-tracker-with-go/db"
	"example.com/expense-tracker-with-go/utils"
)

type User struct {
	ID       int64
	Fullname string `json:"fullname"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string
}

func (u User) SaveNewUser() error {
	query := `INSERT INTO users(fullname, email, password, role) VALUES (?, ?, ?, ?)`

	stmt, err := db.DataBase.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	u.Role = "USER"
	encryptedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Fullname, u.Email, encryptedPass, u.Role)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

func (u *User) ValidatingCredentials() error {
	query := `SELECT id, fullname, password, role FROM users WHERE email = ?`

	row := db.DataBase.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &u.Fullname, &retrievedPassword, &u.Role)
	if err != nil {
		return errors.New("credentials invalid")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	return nil
}
