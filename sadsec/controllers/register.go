package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *Application) Register(c *gin.Context) {
	app.Authorize(c)

	if c.Request.Method == http.MethodPost {
		username := c.PostForm("username")
		email := c.PostForm("email")
		password := c.PostForm("password1")
		password2 := c.PostForm("password2")

		username = strings.TrimSpace(username)
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		password2 = strings.TrimSpace(password2)

		if username == "" || email == "" || password == "" || password2 == "" {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"Gerr": "ALL Fields Required"})
			return
		}

		msg := app.ValidateUsername(username)
		if msg != nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"Uerr": msg})
			return
		}

		check := app.UsernameExists(username)
		if check == true {
			uErr := errors.New("Username Exist")
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"Uerr": uErr})
			return
		}

		err := app.ValidateEmail(email, c)
		valid := true
		if err != valid {
			emErr := errors.New("Enter a valid email")
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"Eerr": emErr})
			return
		}

		pErr := app.CheckP2(password, password2)
		if pErr != nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"Perr": pErr})
			return
		}

		msg = app.ValidatePassword(password)
		if msg != nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"Perr": msg})
			return
		}

		iv, hashpass := app.RegisterHashPassword(password)
		verificationToken := iv

		er := app.Commands.CreateUser(username, email, hashpass, iv, verificationToken)
		if er != nil {
			c.Redirect(302, "/500")
			return
		} else {
			//sendEmail(email, username, email)
			//c.Redirect(301, "/verify_your_email")
			msg := "verification link sent to your email"
			c.HTML(http.StatusOK, "verify_email.html", gin.H{"msg": msg})
			return
		}
	}

	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "register.html", nil)
		return
	}

	c.HTML(http.StatusBadRequest, "register.html", nil)
}

func (app *Application) iv() string {
	seed := time.Now().UTC().UnixNano()
	var Iv = make([]byte, 32)
	rand.Seed(seed)
	rand.Read(Iv)
	res := hex.EncodeToString(Iv)
	res = strings.TrimSpace(res)
	return res
}

func (app *Application) RegisterHashPassword(password string) (string, string) {
	password = strings.TrimSpace(password)

	iv := app.iv()

	outpassword := password + iv
	outpassword = strings.TrimSpace(outpassword)

	hashPassword := sha256.Sum256([]byte(outpassword))
	hashPass := hex.EncodeToString(hashPassword[:])
	pass := strings.TrimSpace(hashPass)

	return iv, pass
}

func sendEmail(reciever, username, email string) {
	to := []string{}
	to = append(to, reciever)

	msg := []byte("verify email http://127.0.0.1:8000/verify_your_email/" + username + "/" + email)

	from := "n33ds0n@gmail.com"
	password := "systemprogemming007"
	smtpHost := "smtp.gmail.com"
	smptPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	smtp.SendMail(smtpHost+":"+smptPort, auth, from, to, msg)
}
