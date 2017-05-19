package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"proxysql-master/admin/cmd"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	//	"proxysql-master/admin/servers"
	"proxysql-master/admin/users"
)

var (
	db  *sql.DB
	err error
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err = sql.Open("mysql", "admin:admin@tcp(172.18.7.204:6032)/main?charset=utf8")
	if err != nil {
		log.Fatal("Open()", err)
	}
	defer db.Close()

	e.GET("/users", listAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:username", listOneUser)
	//	e.PUT("/users/:username", updateUsers)
	e.DELETE("/users/:username", deleteOneUser)

	e.Logger.Fatal(e.Start(":3333"))
}
