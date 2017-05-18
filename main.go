package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"log"
	"proxysql-master/admin/users"
)

func main() {
	db, err := sql.Open("mysql", "admin:admin@tcp(172.18.7.204:6032)/main?charset=utf8")
	if err != nil {
		log.Fatal("Open()", err)
	}
	defer db.Close()

	var proxysql_users users.Users
	proxysql_users.Username = "tianlei2"
	proxysql_users.Password = "111111"

	proxysql_users.AddOneUser(db)
	//proxysql_users.FindOneUserInfo(db)
	users.FindAllUserInfo(db)
	//proxysql_users.DeleteOneUser(db)
}
