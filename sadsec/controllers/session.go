package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

const (
	sessionKey = "3145e33d938cbf1927c2e99f2057586bb85e17fbc2ce884b555cd768641b29a9"
)

var store = sessions.NewCookieStore([]byte(sessionKey))

func (app *Application) CreateSession(c *gin.Context, username string) {

	session, _ := store.Get(c.Request, "session")
	session.Values["userID"] = username
	_ = store.Save(c.Request, c.Writer, session)
}

func (app *Application) CheckSession(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	_, ok := session.Values["userID"]
	if ok {
		c.Next()
	}

	c.Redirect(302, "/login")
	c.Abort()
	return

}

func (app *Application) Authorize(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	_, ok := session.Values["userID"]
	if ok {
		c.Redirect(302, "/")
	}

	return

}

func (app *Application) Logout(c *gin.Context) {
	session, err := store.Get(c.Request, "session")
	if err != nil {
		fmt.Println(err)
	}
	_, ok := session.Values["userID"]

	if ok {
		session.Values["userID"] = ""
		err = store.Save(c.Request, c.Writer, session)
		if err != nil {
			fmt.Println(err)
		}

		app.Login(c)
	}
}
