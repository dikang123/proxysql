package proxysql

import (
	"flag"
	"testing"
)

func TestNewConn(t *testing.T) {

	flag.Parse()
	conn, err := NewConn(*proxysql_addr, *proxysql_port, *proxysql_user, *proxysql_pass)
	if err != nil {
		t.Error(conn, err)
	}

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI()

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}

}
