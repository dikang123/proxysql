package main

import (
	"database/sql"
	//	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	//	"github.com/labstack/echo/middleware"
	"log"
	//	"net/http"
	"proxysql-master/pmapi"
	//	"proxysql-master/admin/servers"
	//"proxysql-master/admin/cmd"
	//	"proxysql-master/admin/users"
)

func main() {
	var err error
	pmapiv1 := new(pmapi.PMApi)

	pmapiv1.Echo = echo.New()
	e := pmapiv1.Echo
	pmapiv1.PMuser = "admin"
	pmapiv1.PMpass = "admin"
	pmapiv1.PMhost = "172.18.7.204:6032"
	pmapiv1.PMdb = "main"
	pmapiv1.MakePMdbi()

	pmapiv1.RegisterMiddleware()

	pmapiv1.Apidb, err = sql.Open("mysql", "admin:admin@tcp(172.18.7.204:6032)/main?charset=utf8")
	if err != nil {
		log.Fatal("Open()", err)
	}
	defer pmapiv1.Apidb.Close()

	e.GET("/users", pmapiv1.ListAllUsers)
	e.POST("/users", pmapiv1.CreateUser)
	e.GET("/users/:username", pmapiv1.ListOneUser)
	//	e.PUT("/users/:username", updateUsers)
	e.DELETE("/users/:username", pmapiv1.DeleteOneUser)

	e.Logger.Fatal(e.Start(":3333"))
}
