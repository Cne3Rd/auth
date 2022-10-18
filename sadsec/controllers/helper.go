package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) InternalServerError(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "500.html", nil)
}
