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

func (u User) CreateAdmin() error {
	query := `INSERT INTO users(fullname, email, password, role) VALUES (?, ?, ?, ?)`

	stmt, err := db.DataBase.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	u.Role = "ADMIN"
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

func (u User) RemoveUserByID() error {
	query := "DELETE FROM users WHERE id = ?"

	stmt, err := db.DataBase.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(usrID int64) (*User, error) {
	query := `SELECT * FROM users WHERE id = ? AND role = ?`

	row := db.DataBase.QueryRow(query, usrID, "USER")
	
	var user User
	err := row.Scan(&user.ID, &user.Fullname, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsersData() ([]User, error) {
	query := `SELECT * FROM users WHERE role = ?`

	rows, err := db.DataBase.Query(query, "USER")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allUser []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}

		allUser = append(allUser, user)
	}

	return allUser, nil
}
