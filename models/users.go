package models

import (
	"errors"

	"github.com/warlock1729/first-go-project/db"
	"github.com/warlock1729/first-go-project/utils"
)

type User struct {
	ID       int64
	Name     string 
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
		INSERT
		INTO users (name,email,password)
		values (?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Name, user.Email, passwordHash)
	if err != nil {
		return err
	}
	user_id, err := result.LastInsertId()
	user.ID = user_id
	defer stmt.Close()
	return nil
}

func (user *User) ValidateCredentials() error {
	query := `
	select id,password from users where email = ?
	`
	row := db.DB.QueryRow(query, user.Email)

	var (
		retrievedPasswordHash string
	)

	// the args are passed in order of the selected columns in query
	err := row.Scan(&user.ID, &retrievedPasswordHash)

	if err != nil {
		return errors.New("Invalid credentials")
	}
	isPasswordMatch := utils.CheckPasswordHash(user.Password, retrievedPasswordHash)

	if !isPasswordMatch {
		return errors.New("Invalid credentials")
	}
	return nil

}
