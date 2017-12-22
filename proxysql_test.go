package proxysql

import "testing"

func TestNewConn(t *testing.T) {
	conn, err := NewConn("172.18.10.111", 13306, "admin", "admin")
	if err != nil {
		t.Error(conn, err)
	}

	db, err := conn.OpenConn()
	if err != nil {
		t.Error(db, err)
	}
}
