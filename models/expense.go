package models

import (
	"time"

	"example.com/expense-tracker-with-go/db"
)

type Expense struct {
	ID           int64
	Account      string    `binding:"required"`
	Amount       int64     `binding:"required"`
	Category     string    `binding:"required"`
	Date         time.Time `binding:"required"`
	Expense_type string    `binding:"required"`
	Note         string    `binding:"required"`
	User_ID      int64
}

func (e *Expense) SaveExpense() error {
	query := `INSERT INTO expenses(account, amount, category, date, expense_type, note, user_id)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	stmt, err := db.DataBase.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Account, e.Amount, e.Category, e.Date, e.Expense_type, e.Note, e.User_ID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return nil
}

func (e Expense) UpdateExpenseByID() error {
	query := `UPDATE expenses SET account = ?, amount = ?, category = ?, date = ?, expense_type = ?, note = ?
	WHERE id = ?`

	stmt, err := db.DataBase.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(e.Account, e.Amount, e.Category, e.Date, e.Expense_type, e.Note, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetExpensebyID(eID int64) (*Expense, error) {
	query := `SELECT * FROM expenses WHERE id = ?`

	row := db.DataBase.QueryRow(query, eID)

	var expense Expense
	err := row.Scan(
		&expense.ID, &expense.Account, &expense.Amount, &expense.Category, &expense.Date,
		&expense.Expense_type, &expense.Note, &expense.User_ID)
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func GetAllExpensesByUserID(uID int64) ([]Expense, error) {
	query := "SELECT * FROM expenses WHERE user_id = ?"

	rows, err := db.DataBase.Query(query, uID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var allExpenses []Expense
	for rows.Next() {
		var expense Expense
		err := rows.Scan(
			&expense.ID, &expense.Account, &expense.Amount, &expense.Category, &expense.Date,
			&expense.Expense_type, &expense.Note, &expense.User_ID)
		if err != nil {
			return nil, err
		}

		allExpenses = append(allExpenses, expense)
	}

	return allExpenses, nil
}

func (e Expense) DeleteExpenseByID() error {
	query := `DELETE FROM expenses WHERE id = ?`

	stmt, err := db.DataBase.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}

	return nil
}
