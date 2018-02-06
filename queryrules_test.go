package proxysql

import (
	"flag"
	"testing"
)

func TestFindAllQr(t *testing.T) {

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

	allqr, err := FindAllQr(db, 1, 0)
	if err != nil {
		t.Error(allqr, err)
	}

	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}

}
func TestAddOneQr(t *testing.T) {

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

	newqr, err := NewQr("dev", 0)
	if err != nil {
		t.Error(err)
	}

	err = newqr.AddOneQr(db)
	if err != nil {
		t.Error(err)
	}

	newqr.SetQrActive(1)
	newqr.SetQrMatchDigest("^SELECT")

	err = newqr.UpdateOneQrInfo(db)
	if err != nil {
		t.Error(err)
	}

	err = newqr.DeleteOneQr(db)
	if err != nil {
		t.Error(err)
	}

	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}
}
