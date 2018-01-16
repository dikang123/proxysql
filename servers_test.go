package proxysql

import (
	"testing"
)

func TestFindAllServers(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
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

	allservers, err := FindAllServerInfo(db, 1, 0)
	if err != nil {
		t.Error(allservers, err)
	}

}

func TestAddOneServer(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
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

	newsrv, err := NewServer(1, "192.168.100.111", 6032)

	err = newsrv.AddOneServers(db)
	if err != nil {
		t.Error(err)
	}

}

func TestUpdateOneServer(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
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

	newsrv, err := NewServer(1, "192.168.100.111", 6032)
	newsrv.SetServerMaxConnection(9999)
	newsrv.SetServersComment("test hostgroup")
	newsrv.SetServerStatus("ONLINE")

	err = newsrv.UpdateOneServerInfo(db)
	if err != nil {
		t.Error(err)
	}

}

func TestDeleteOneServer(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
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

	newsrv, err := NewServer(1, "192.168.100.111", 6032)

	err = newsrv.DeleteOneServers(db)
	if err != nil {
		t.Error(err)
	}

}
