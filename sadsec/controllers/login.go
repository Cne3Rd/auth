package controllers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Logins struct {
	username string
	password string
	iv       string
}

func (app *Application) Login(c *gin.Context) {
	app.Authorize(c)

	// check if http request verb is POST
	// if it is perform operations
	if c.Request.Method == http.MethodPost {
		username := c.PostForm("username")
		password := c.PostForm("password1")

		username = strings.TrimSpace(username)
		password = strings.TrimSpace(password)

		// check if any of user input does is empty
		// if it is perform oprations
		if username == "" || password == "" {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": "All Fields Required"})
			return
		}

		// query the database to retrieve data needed to perform operation
		row := app.Commands.LoginUser(username)

		// initialize struct
		// will be holding data retrieve from the database
		user := &Logins{}

		// stored data retrieve from the database into struct field
		err := row.Scan(&user.username, &user.password, &user.iv)
		if err == sql.ErrNoRows {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": "Username or Password Incorrect"})
			return
		}

		// hashing the user input password with the user iv save in the database
		password = app.LoginHashPassword(password, user.iv)

		// comparing user input data to the one save in the database
		if user.username == username && password == user.password {
			re := app.IsActive(username)
			if re == false {
				c.HTML(http.StatusForbidden, "must_verify.html", nil)
				return
			}
			app.CreateSession(c, user.username)
			c.Redirect(302, "/")
			return
		}

		c.HTML(http.StatusOK, "login.html", gin.H{"err": "Username or Password Incorrect"})

	}

	// check if http request verb is GET
	// if it is perform oprations
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}

}

func (app *Application) LoginHashPassword(password, iv string) string {
	password = strings.TrimSpace(password)
	iv = strings.TrimSpace(iv)

	outpassword := password + iv
	outpassword = strings.TrimSpace(outpassword)

	hashPassword := sha256.Sum256([]byte(outpassword))
	hashPass := hex.EncodeToString(hashPassword[:])
	pass := strings.TrimSpace(hashPass)

	return pass
}
