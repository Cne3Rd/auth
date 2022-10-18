package models

import (
	"database/sql"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) CreateUser(username, email, hashpass, iv, vtoken string) error {
	stmt := "INSERT INTO users(username, email, password, iv, verification_token) VALUES($1, $2, $3, $4, $5)"
	exec, err := db.DB.Exec(stmt, username, email, hashpass, iv, vtoken)
	if err != nil {
		return err
	}

	_ = exec
	return nil
}

func (db *Database) LoginUser(username string) *sql.Row {
	stmt := "SELECT username, password, iv FROM users WHERE username=$1"
	row := db.DB.QueryRow(stmt, username)
	return row
}

func (db *Database) GetUid(username string) *sql.Row {
	stmt := "SELECT id FROM users WHERE username = $1"
	row := db.DB.QueryRow(stmt, username)
	return row
}

func (db *Database) GetVToken(username string) *sql.Row {
	stmt := "SELECT verification_token FROM users where username=$1"
	row := db.DB.QueryRow(stmt, username)
	return row
}

func (db *Database) IsActive(username string) *sql.Row {
	stmt := "SELECT is_active FROM users WHERE username=$1"
	row := db.DB.QueryRow(stmt, username)
	return row
}
