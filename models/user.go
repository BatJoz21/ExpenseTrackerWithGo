package models

import (
	"example.com/expense-tracker-with-go/db"
	"example.com/expense-tracker-with-go/utils"
)

type User struct {
	ID       int64
	Fullname string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
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
