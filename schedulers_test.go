package proxysql

import (
	"testing"
)

func TestFindAllSchedulers(t *testing.T) {
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

	allservers, err := FindAllSchedulerInfo(db, 1, 0)
	if err != nil {
		t.Error(allservers, err)
	}

}

func TestAddOneSchedulers(t *testing.T) {
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

	sched, err := NewSch("/bin/bash", 1000)
	if err != nil {
		t.Error(err)
	}

	err = sched.AddOneScheduler(db)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateOneSchedulers(t *testing.T) {
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

	sched, err := NewSch("/bin/bash", 1000)
	if err != nil {
		t.Error(err)
	}

	sched.Id = 2
	sched.Active = 1
	sched.Arg1 = "ls -l"

	err = sched.UpdateOneSchedulerInfo(db)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteOneSchedulers(t *testing.T) {
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

	sched, err := NewSch("/bin/bash", 1000)
	if err != nil {
		t.Error(err)
	}

	sched.Id = 1

	err = sched.DeleteOneScheduler(db)
	if err != nil {
		t.Error(err)
	}
}
