package proxysql

import "testing"

func TestNewConn(t *testing.T) {
	conn, err := NewConn("172.18.10.136", 13306, "admin", "admin")
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
