package main

import (
	"log"
	"proxysql-master/admin/servers"
	"upper.io/db.v3/mysql"
)

var settings = mysql.ConnectionURL{
	Host:     "172.18.7.204:6032",
	User:     "admin",
	Password: "admin",
	Database: "stats",
	Options:  map[string]string{"parseTime": "false"},
}

func main() {
	sess, err := mysql.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	defer sess.Close()
	var proxysql_mysql_servers []servers.Servers

	err = sess.Collection("mysql_servers").Find().All(&proxysql_mysql_servers)
	if err != nil {
		log.Fatalf("Find(): %q\n", err)
	}

	for _, srvs := range proxysql_mysql_servers {
		log.Printf("%#v\n", srvs)
	}
}
