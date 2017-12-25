package proxysql

import (
	"testing"
)

func TestGetConfigs(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	allusers, err := GetConfig(db)
	if err != nil {
		t.Error(allusers, err)
	}

}

func TestUpdateOneConfigs(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}

	err = UpdateOneConfig(db, "mysql-max_connections", "99999")
	if err != nil {
		t.Error(err)
	}

}