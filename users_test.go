package proxysql

import (
	"flag"
	"testing"
)

var proxysql_addr = flag.String("addr", "127.0.0.1", "proxysql listen address.default 127.0.0.1")
var proxysql_port = flag.Uint64("port", 6032, "proxysql listen port,default 6032")
var proxysql_user = flag.String("user", "admin", "proxysql administrator name.default admin")
var proxysql_pass = flag.String("pass", "admin", "proxysql administrator password.default admin")

func TestFindAllUsers(t *testing.T) {

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

	allusers, err := FindAllUserInfo(db, 1, 0)
	if err != nil {
		t.Error(allusers, err)
	}

	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}

}

func TestAddOneUser(t *testing.T) {

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

	newuser, err := NewUser("devtest", "devtest", 0, "dev")
	if err != nil {
		t.Error(err)
	}
	newuser.SetUserActive(1)

	err = newuser.AddOneUser(db)
	if err != nil {
		t.Error(err)
	}

	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}

}

func TestUpdateOneUser(t *testing.T) {

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

	newuser, err := NewUser("devtest", "devtest", 0, "dev")
	if err != nil {
		t.Error(err)
	}
	newuser.SetMaxConnections(999)
	newuser.SetSchemaLocked(1)

	err = newuser.UpdateOneUserInfo(db)
	if err != nil {
		t.Error(err)
	}

	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}

}

func TestDeleteOneUser(t *testing.T) {

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

	newuser, err := NewUser("devtest", "devtest", 0, "dev")
	if err != nil {
		t.Error(err)
	}

	err = newuser.DeleteOneUser(db)
	if err != nil {
		t.Error(err)
	}

	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}

}
