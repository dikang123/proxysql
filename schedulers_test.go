package proxysql

import (
	"flag"
	"testing"
)

func TestFindAllSchedulers(t *testing.T) {

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

	allservers, err := FindAllSchedulerInfo(db, 1, 0)
	if err != nil {
		t.Error(allservers, err)
	}
	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}

}

func TestAddOneSchedulers(t *testing.T) {

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

	sched, err := NewSch("/bin/bash", 1000)
	if err != nil {
		t.Error(err)
	}

	err = sched.AddOneScheduler(db)
	if err != nil {
		t.Error(err)
	}

	sched.Active = 1
	sched.Arg1 = "ls -l"

	err = sched.UpdateOneSchedulerInfo(db)
	if err != nil {
		t.Error(err)
	}

	err = sched.DeleteOneScheduler(db)
	if err != nil {
		t.Error(err)
	}

	err = conn.CloseConn(db)
	if err != nil {
		t.Error(err)
	}
}
