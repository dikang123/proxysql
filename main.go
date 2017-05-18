package main

import (
	//	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//	"github.com/pressly/chi"
	//	"github.com/pressly/chi/middleware"
	"log"
	"proxysql-master/admin/servers"
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

	var proxysql_servers servers.Servers
	proxysql_servers.HostGroupId = 3
	proxysql_servers.HostName = "dn10"
	proxysql_servers.Port = 3310

	proxysql_users.AddOneUser(db)
	//proxysql_users.FindOneUserInfo(db)
	//users.FindAllUserInfo(db)
	//proxysql_users.DeleteOneUser(db)

	proxysql_servers.AddOneServers(db)
	//proxysql_servers.FindOneServersInfo(db)
	proxysql_servers.DeleteOneServers(db)
	//servers.FindAllServerInfo(db)
}
