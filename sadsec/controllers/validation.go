package controllers

import (
	"database/sql"
	"errors"
	"unicode"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
	Vpassword "github.com/wagslane/go-password-validator"
)

func (app *Application) ValidateUsername(user string) error {
	for _, char := range user {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return errors.New("Only alphanumeric characters allowed for username ")
		}
	}
	if len(user) >= 6 {
		return nil
	} else {
		return errors.New("username length must be greater than 5")
	}

}

func (app *Application) UsernameExists(username string) bool {
	exists := true

	row := app.Commands.GetUid(username)
	var uid int

	err := row.Scan(&uid)
	if err == sql.ErrNoRows {
		return false
	}

	return exists

}

func (app *Application) ValidateEmail(email string, c *gin.Context) bool {
	valid := true

	check := emailverifier.IsAddressValid(email)
	if check != valid {
		return false
	}

	return valid

}

func (app *Application) ValidatePassword(pass string) error {
	minENtropyBits := 60.0
	err := Vpassword.Validate(pass, minENtropyBits)
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) CheckP2(pass1, pass2 string) error {
	if pass1 != pass2 {
		return errors.New("Password Not Match")
	}

	return nil
}

func (app *Application) IsActive(username string) bool {
	var active bool
	stmt := "SELECT is_active from users WHERE username=$1"
	row := app.Commands.DB.QueryRow(stmt, username)
	err := row.Scan(&active)
	_ = err

	if active != true {
		return false
	}

	return true

}
