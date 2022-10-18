package controllers

import (
	"log"
	"net/http"

	"github.com/cne3rd/sadsec/models"
	"github.com/gin-gonic/gin"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Commands *models.Database
}

func (app *Application) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
