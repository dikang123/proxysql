// proxysql

package proxysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/juju/errors"
)

type (
	// connect proxysql use admin user.
	Conn struct {
		Addr      string
		Port      uint64
		User      string
		Password  string
		Database  string
		Charset   string
		Collation string
		DBI       string
		Retry     uint64
	}
)

func NewConn(addr string, port uint64, user string, password string) (*Conn, error) {

	ps := new(Conn)
	ps.Addr = addr
	ps.Port = port
	ps.User = user
	ps.Password = password
	ps.Database = "stats"
	ps.Charset = "utf8"
	ps.Collation = "utf8_general_ci"
	ps.Retry = 3

	return ps, nil
}

// set character set .such as : utf8
func (ps *Conn) SetCharset(charset string) {
	ps.Charset = charset
}

// set collation.such as : utf8_general_ci
func (ps *Conn) SetCollation(collation string) {
	ps.Collation = collation
}

// set retrys.
func (ps *Conn) SetRetry(retry uint64) {
	ps.Retry = retry
}

func (ps *Conn) MakeDBI() {
	ps.DBI = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s", ps.User, ps.Password, ps.Addr, ps.Port, ps.Database, ps.Charset, ps.Collation)
}

func (ps *Conn) OpenConn() (*sql.DB, error) {

	db, err := sql.Open("mysql", ps.DBI)
	if err != nil {
		return nil, errors.Trace(err)
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Trace(err)
	}
	//defer db.Close()

	return db, nil
}

// close connection.
func (ps *Conn) CloseConn(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
