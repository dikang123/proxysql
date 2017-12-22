package proxysql

import (
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	allusers, err := FindAllUserInfo(db, 1, 0)
	if err != nil {
		t.Error(allusers, err)
	}

}

func TestAddOneUser(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	newuser := new(Users)
	newuser.SetDefaultHostgroup(0)
	newuser.SetDefaultSchema("dev")
	newuser.SetUserName("devtest")
	newuser.SetUserPass("devtest")

	err = newuser.AddOneUser(db)
	if err != nil {
		t.Error(err)
	}

}

func TestDeleteOneUser(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	newuser := new(Users)
	newuser.SetUserName("devtest")
	newuser.SetBackend(1)
	newuser.SetFrontend(1)

	err = newuser.DeleteOneUser(db)
	if err != nil {
		t.Error(err)
	}

}
