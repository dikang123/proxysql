// proxysql

package proxysql

import (
	"database/sql"
	"fmt"
)

type (
	// connect proxysql use admin user.
	Conn struct {
		Addr      string
		Port      int
		User      string
		Password  string
		Database  string
		Charset   string
		Collation string
		DBI       string
	}
)

func NewConn(addr string, port int, user string, password string) (*Conn, error) {
	ps := new(Conn)
	ps.Addr = addr
	ps.Port = port
	ps.User = user
	ps.Password = password
	ps.Database = "stats"
	ps.Charset = "utf8"
	ps.Collation = "utf8_general_ci"

	return ps, nil
}

func (ps *Conn) SetCharset(charset string) {
	ps.Charset = charset
}

func (ps *Conn) SetCollation(collation string) {
	ps.Collation = collation
}

func (ps *Conn) OpenConn() (*sql.DB, error) {
	ps.DBI = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8?collation=utf8_general_ci", ps.User, ps.Password, ps.Addr, ps.Port, ps.Database)

	db, err := sql.Open("mysql", ps.DBI)
	if err != nil {
		println(err)
	}

	err = db.Ping()
	if err != nil {
		println(err)
	}
	return db, nil
}
