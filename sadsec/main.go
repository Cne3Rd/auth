package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	controller "github.com/cne3rd/sadsec/controllers"
	"github.com/cne3rd/sadsec/models"
	router "github.com/cne3rd/sadsec/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Database connection
	dbconn, err := DatabaseConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer dbconn.Close()

	db := &models.Database{DB: dbconn}

	errLog := log.New(os.Stderr, "ERROR	\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &controller.Application{
		ErrorLog: errLog,
		InfoLog:  infoLog,
		Commands: db,
	}

	r := gin.Default()

	// set templates

	r.LoadHTMLGlob("templates/*")

	// set static files

	r.StaticFS("/static", http.Dir("./static"))

	rout := router.Control{
		Route: r,
		App:   app,
	}

	rout.Routes()

	srv := &http.Server{
		Addr:     "127.0.0.1:8000",
		ErrorLog: app.ErrorLog,
		Handler:  rout.Route,
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func DatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "sad.db")
	if err != nil {
		fmt.Println("Database connections failed:", err)
		return nil, err
	}

	perr := db.Ping()
	if perr != nil {
		fmt.Println("Can't reach database:", perr)
		return nil, perr
	}

	return db, nil

}
