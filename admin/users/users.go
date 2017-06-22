package users

/*
* 1.获取当前用户列表
* 2.通过用户名获取用户的详细信息
* 3.添加一个新用户
* 4.删除一个用户
* 5.激活一个用户
* 6.反激活一个用户
* 7.更新用户信息
*
 */

import (
	"database/sql"
	"fmt"
	"log"
	"proxysql-master/admin/cmd"
	//	"os"
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
	StmtUserExist         = `SELECT count(*) FROM mysql_users WHERE username = %q`
	StmtAddOneUser        = `INSERT INTO mysql_users(username,password) VALUES(%q,%q)`
	StmtDeleteOneUser     = `DELETE FROM mysql_users WHERE username = %q`
	StmtActiveOneUser     = `UPDATE mysql_users SET active = 1 WHERE username = %q`
	StmtDisactiveOneUser  = `UPDATE mysql_users SET active = 0 WHERE username = %q`
	StmtFindOneUserInfo   = `SELECT ifnull(username,""),ifnull(password,""),ifnull(active,0),ifnull(use_ssl,0),ifnull(default_hostgroup,0),ifnull(default_schema,""),ifnull(schema_locked,0),ifnull(transaction_persistent,0),ifnull(fast_forward,0),ifnull(backend,0),ifnull(frontend,0),ifnull(max_connections,0) FROM mysql_users WHERE username = %q`
	StmtFindAllUserInfo   = `SELECT ifnull(username,""),ifnull(password,""),ifnull(active,0),ifnull(use_ssl,0),ifnull(default_hostgroup,0),ifnull(default_schema,""),ifnull(schema_locked,0),ifnull(transaction_persistent,0),ifnull(fast_forward,0),ifnull(backend,0),ifnull(frontend,0),ifnull(max_connections,0) FROM mysql_users limit %d offset %d`
	StmtUpdateOneUserDs   = `UPDATE mysql_users SET default_schema=%q WHERE username = %q`
	StmtUpdateOneUserMc   = `UPDATE mysql_users SET max_connections = %d WHERE username = %q`
	StmtUpdateOneUserDH   = `UPDATE mysql_users SET default_hostgroup=%d WHERE username = %q`
	StmtUpdateOneUserPass = `UPDATE mysql_users SET password=%q WHERE username = %q`
)

func (users *Users) UserExists(db *sql.DB) bool {
	st := fmt.Sprintf(StmtUserExist, users.Username)
	rows, err := db.Query(st)
	if err != nil {
		log.Print("UserExists:", err)
	}
	var UserCount uint64
	for rows.Next() {
		err = rows.Scan(&UserCount)
		if err != nil {
			log.Print("UserExists,Scan:", err)
		}
	}
	if UserCount == 0 {
		return false
	} else {
		return true
	}
}

func (users *Users) AddOneUser(db *sql.DB) int {
	if isexist := users.UserExists(db); isexist == false {
		st := fmt.Sprintf(StmtAddOneUser, users.Username, users.Password)
		_, err := db.Query(st)
		if err != nil {
			return 1 //add user failed
		}
		cmd.LoadUserToRuntime(db)
		cmd.SaveUserToDisk(db)
		return 0
	} else {
		return 2 //username exists
	}
}

func (users *Users) DeleteOneUser(db *sql.DB) int {
	if isexist := users.UserExists(db); isexist == true {
		st := fmt.Sprintf(StmtDeleteOneUser, users.Username)
		_, err := db.Query(st)
		if err != nil {
			return 1 //delte failed
		}
		cmd.LoadUserToRuntime(db)
		cmd.SaveUserToDisk(db)
		return 0 //delete success

	} else {
		return 2 //user exists
	}
}

func (users *Users) ActiveOneUser(db *sql.DB) int {
	if isexist := users.UserExists(db); isexist == true {
		st := fmt.Sprintf(StmtActiveOneUser, users.Username)
		_, err := db.Query(st)
		if err != nil {
			log.Print("ActiveOneUser:", err)
			return 1
		}
		cmd.LoadUserToRuntime(db)
		cmd.SaveUserToDisk(db)
		return 0
	} else {
		log.Print("ActiveOneUser: User is not exists")
		return 2
	}
}

func (users *Users) DisactiveOneUser(db *sql.DB) int {
	if isexist := users.UserExists(db); isexist == true {
		st := fmt.Sprintf(StmtDisactiveOneUser, users.Username)
		_, err := db.Query(st)
		if err != nil {
			//log.Print("DisactiveOneUser:", err)
			return 1
		}
		cmd.LoadUserToRuntime(db)
		cmd.SaveUserToDisk(db)
		return 0
	} else {
		//log.Print("DisactiveOneUser: User is not exists")
		return 2
	}
}

func (users *Users) UpdateOneUserDh(db *sql.DB) int {
	if isexist := users.UserExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneUserDH, users.DefaultHostgroup, users.Username)
		_, err := db.Query(st)
		if err != nil {
			log.Print("UpdateOneUserDH()", err)
			return 1
		}
		cmd.LoadUserToRuntime(db)
		cmd.SaveUserToDisk(db)
		return 0
	} else {
		log.Print("UpdateOneUserDH()", "User is not exists")
		return 2
	}
}

func (users *Users) UpdateOneUserPass(db *sql.DB) int {
	if isexist := users.UserExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneUserPass, users.Password, users.Username)
		_, err := db.Query(st)
		if err != nil {
			log.Print("UpdateOneUserPass()", err)
			return 1
		}
		cmd.LoadUserToRuntime(db)
		cmd.SaveUserToDisk(db)
		return 0
	} else {
		log.Print("UpdateOneUserPass()", "User is not exists")
		return 2
	}
}

func (users *Users) UpdateOneUserDs(db *sql.DB) int {
	if isexist := users.UserExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneUserDs, users.DefaultSchema, users.Username)
		_, err := db.Query(st)
		if err != nil {
			log.Print("UpdateOneUserDs:", err)
			return 1
		}
		cmd.LoadUserToRuntime(db)
		cmd.SaveUserToDisk(db)
		return 0
	} else {
		log.Print("UpdateOneUserDs: User is not exists")
		return 2
	}
}

func (users *Users) UpdateOneUserMc(db *sql.DB) int {
	if isexist := users.UserExists(db); isexist == true {
		st := fmt.Sprintf(StmtUpdateOneUserMc, users.MaxConnections, users.Username)
		_, err := db.Query(st)
		if err != nil {
			log.Print("UpdateOneUserMc:", err)
			return 1
		}
		cmd.LoadUserToRuntime(db)
		cmd.SaveUserToDisk(db)
		return 0
	} else {
		log.Print("UpdateOneUserMc: User is not exists")
		return 2
	}
}

func (users *Users) FindOneUserInfo(db *sql.DB) Users {
	if isexist := users.UserExists(db); isexist == true {
		st := fmt.Sprintf(StmtFindOneUserInfo, users.Username)
		rows, err := db.Query(st)
		if err != nil {
			log.Print("UpdateOneUserMc:", err)
		}
		for rows.Next() {
			err = rows.Scan(&users.Username,
				&users.Password,
				&users.Active,
				&users.UseSsl,
				&users.DefaultHostgroup,
				&users.DefaultSchema,
				&users.SchemaLocked,
				&users.TransactionPersistent,
				&users.FastForward,
				&users.Backend,
				&users.Frontend,
				&users.MaxConnections)
		}
	} else {
		log.Print("UpdateOneUserMc: User is not exists")
	}
	return *users
}

func FindAllUserInfo(db *sql.DB, limit int64, skip int64) []Users {
	var alluser []Users
	var tmpusr Users
	var QueryText string
	QueryText = fmt.Sprintf(StmtFindAllUserInfo, limit, skip)
	log.Printf(StmtFindAllUserInfo, limit, skip)
	rows, err := db.Query(QueryText)
	if err != nil {
		log.Print("FindAllUserInfo:", err)
	}
	defer rows.Close()
	for rows.Next() {
		tmpusr = Users{}
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
		alluser = append(alluser, tmpusr)
	}
	return alluser
}
