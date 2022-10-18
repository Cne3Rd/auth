package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) VerifyEmail(c *gin.Context) {

	username := c.Param("username")
	verificationToken := c.Param("verify")

	var active bool
	/*
		row := app.Commands.IsActive(username)
		err := row.Scan(&active)
		if err == sql.ErrNoRows {
			c.HTML(http.StatusInternalServerError, "500.html", nil)
			return
		}
	*/

	active = app.IsActive(username)

	msg := "Verification code invalid or you already been verified    "
	if active == true {
		c.HTML(http.StatusOK, "done_verification.html", gin.H{"msg": msg})
		return
	}

	var dbtoken string
	row := app.Commands.GetVToken(username)
	err := row.Scan(&dbtoken)
	if err == sql.ErrNoRows {
		c.HTML(http.StatusUnauthorized, "register.html", gin.H{"verr": "pls try your verification link again"})
		return
	}

	if verificationToken == dbtoken {
		active := true
		stmt := "UPDATE users SET is_active = $1 WHERE username=$2"
		exec, err := app.Commands.DB.Exec(stmt, active, username)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "500.html", nil)
		}

		_ = exec

		msg := "<h3> Thank You!</h3> <br><br> We have verified your email.<br><br>"
		c.HTML(http.StatusOK, "email_verified.html", gin.H{"msg": msg})
		return
	} else {
		msg := "unable to verify your emial try again"
		c.HTML(http.StatusBadRequest, "bad_verification.html", gin.H{"msg": msg})
		return
	}
}
