package proxysql

import (
	"database/sql"
	"fmt"

	"github.com/juju/errors"
)

type (
	Users struct {
		Username              string `db:"username" json:"username"`
		Password              string `db:"password" json:"password"`
		Active                uint64 `db:"active" json:"active"`
		UseSsl                uint64 `db:"use_ssl" json:"use_ssl"`
		DefaultHostgroup      uint64 `db:"default_hostgroup" json:"default_hostgroup"`
		DefaultSchema         string `db:"default_schema" json:"default_schema"`
		SchemaLocked          uint64 `db:"schema_locked" json:"schema_locked"`
		TransactionPersistent uint64 `db:"transaction_persistent" json:"transaction_persistent"`
		FastForward           uint64 `db:"fast_forward" json:"fast_forward"`
		Backend               uint64 `db:"backend" json:"backend"`
		Frontend              uint64 `db:"frontend" json:"frontend"`
		MaxConnections        uint64 `db:"max_connections" json:"max_connections"`
	}
)

const (
	/*add a new users*/
	StmtAddOneUser = `
	INSERT INTO 
		mysql_users(username,password,default_hostgroup,default_schema)
	VALUES(%q,%q,%d,%q)`

	/*delete a user*/
	StmtDeleteOneUser = `
	DELETE FROM 
		mysql_users 
	WHERE 
		username = %q
	AND
		backend = %d
	AND
		frontend = %d
	`

	/*list all users*/
	StmtFindAllUserInfo = `
	SELECT 
		ifnull(username,""),
		ifnull(password,""),
		ifnull(active,0),
		ifnull(use_ssl,0),
		ifnull(default_hostgroup,0),
		ifnull(default_schema,""),
		ifnull(schema_locked,0),
		ifnull(transaction_persistent,0),
		ifnull(fast_forward,0),
		ifnull(backend,0),
		ifnull(frontend,0),
		ifnull(max_connections,0) 
	FROM mysql_users 
	LIMIT %d 
	OFFSET %d`

	/*update a users*/
	StmtUpdateOneUser = `
	UPDATE 
		mysql_users 
	SET 
		password=%q,
		active=%d,
		use_ssl=%d,
		default_hostgroup=%d,
		default_schema=%q,
		schema_locked=%d,
		transaction_persistent=%d,
		fast_forward=%d,
		backend=%d,
		frontend=%d,
		max_connections=%d 
	WHERE 
		username = %q
	AND
		backend = %d
	AND
		frontend = %d
		`
)

//list all users.
func FindAllUserInfo(db *sql.DB, limit uint64, skip uint64) ([]Users, error) {
	var alluser []Users

	Query := fmt.Sprintf(StmtFindAllUserInfo, limit, skip)

	rows, err := db.Query(Query)
	if err != nil {
		return []Users{}, errors.Trace(err)
	}
	defer rows.Close()

	for rows.Next() {

		var tmpusr Users

		err = rows.Scan(
			&tmpusr.Username,
			&tmpusr.Password,
			&tmpusr.Active,
			&tmpusr.UseSsl,
			&tmpusr.DefaultHostgroup,
			&tmpusr.DefaultSchema,
			&tmpusr.SchemaLocked,
			&tmpusr.TransactionPersistent,
			&tmpusr.FastForward,
			&tmpusr.Backend,
			&tmpusr.Frontend,
			&tmpusr.MaxConnections,
		)

		if err != nil {
			continue
		}

		alluser = append(alluser, tmpusr)
	}
	return alluser, nil
}

func NewUser(username string, password string, default_hostgroup uint64, default_schema string) (*Users, error) {
	newuser := new(Users)

	newuser.Username = username
	newuser.Password = password
	newuser.DefaultHostgroup = default_hostgroup
	newuser.DefaultSchema = default_schema

	return newuser, nil
}

// set fast_forward
func (users *Users) SetFastForward(fast_forward uint64) {
	if fast_forward >= 1 {
		users.FastForward = 1
	} else {
		users.FastForward = 0
	}
}

// set max_connections
func (users *Users) SetMaxConnections(max_connections uint64) {
	switch {
	case max_connections >= 10000:
		users.MaxConnections = 10000
	case max_connections <= 1:
		users.MaxConnections = 1
	default:
		users.MaxConnections = max_connections
	}
}

// set backend
func (users *Users) SetBackend(backend uint64) {
	if backend >= 1 {
		users.Backend = 1
	} else {
		users.Backend = 0
	}
}

// set fronted
func (users *Users) SetFrontend(frontend uint64) {
	if frontend >= 1 {
		users.Frontend = 1
	} else {
		users.Frontend = 0
	}
}

func (users *Users) SetUserActive(active uint64) {
	if active >= 1 {
		users.Active = 1
	} else {
		users.Active = 0
	}
}

//add a new user.
func (users *Users) AddOneUser(db *sql.DB) error {

	Query := fmt.Sprintf(StmtAddOneUser, users.Username, users.Password, users.DefaultHostgroup, users.DefaultSchema)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err) //add user failed
	}

	LoadUserToRuntime(db)
	SaveUserToDisk(db)

	return nil
}

//delete a user.
func (users *Users) DeleteOneUser(db *sql.DB) error {

	Query := fmt.Sprintf(StmtDeleteOneUser, users.Username, users.Backend, users.Frontend)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err) //delte failed
	}

	LoadUserToRuntime(db)
	SaveUserToDisk(db)

	return nil //delete success
}

// update a user.
func (users *Users) UpdateOneUserInfo(db *sql.DB) error {

	Query := fmt.Sprintf(StmtUpdateOneUser,
		users.Password,
		users.Active,
		users.UseSsl,
		users.DefaultHostgroup,
		users.DefaultSchema,
		users.SchemaLocked,
		users.TransactionPersistent,
		users.FastForward,
		users.Backend,
		users.Frontend,
		users.MaxConnections,
		users.Username,
		users.Backend,
		users.Frontend)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	LoadUserToRuntime(db)
	SaveUserToDisk(db)

	return nil
}
