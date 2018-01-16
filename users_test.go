package proxysql

import (
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI

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

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI

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

}

func TestDeleteOneUser(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}
	
	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI

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

}
func TestUpdateOneUser(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	conn.SetCharset("utf8")
	conn.SetCollation("utf8_general_ci")
	conn.MakeDBI

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

}
