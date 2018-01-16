package proxysql

import (
	"testing"
)

func TestFindAllQr(t *testing.T) {
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

	allqr, err := FindAllQr(db, 1, 0)
	if err != nil {
		t.Error(allqr, err)
	}

}
func TestAddOneQr(t *testing.T) {
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

	newqr, err := NewQr("dev", 0)
	if err != nil {
		t.Error(err)
	}

	err = newqr.AddOneQr(db)
	if err != nil {
		t.Error(err)
	}

}

func TestUpdateOneQr(t *testing.T) {
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

	newqr, err := NewQr("dev", 0)
	if err != nil {
		t.Error(err)
	}

	newqr.SetQrRuleid(4)
	newqr.SetQrActive(1)
	newqr.SetQrMatchDigest("^SELECT")

	err = newqr.UpdateOneQrInfo(db)
	if err != nil {
		t.Error(err)
	}

}

func TestDeleteOneQr(t *testing.T) {
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

	newqr, err := NewQr("dev", 0)
	if err != nil {
		t.Error(err)
	}

	newqr.SetQrRuleid(2)

	err = newqr.DeleteOneQr(db)
	if err != nil {
		t.Error(err)
	}

}
