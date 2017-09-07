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
	/*新建一个用户*/
	StmtAddOneUser = `
	INSERT INTO 
		mysql_users(username,password,default_schema) 
	VALUES(%q,%q,%q)`

	/*删除一个用户*/
	StmtDeleteOneUser = `
	DELETE FROM 
		mysql_users 
	WHERE 
		username = %q
	AND
		backend = 1
	AND
		frontend = 1
	`

	/*查询出所有用户的信息*/
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

	/*更新一个用户信息*/
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
		backend = 1
	AND
		frontend = 1
		`
)

func (users *Users) FindAllUserInfo(db *sql.DB, limit int64, skip int64) ([]Users, error) {
	/*定义一个新的变量，alluser用户保存所有用户信息*/
	var alluser []Users

	Query := fmt.Sprintf(StmtFindAllUserInfo, limit, skip)
	log.Printf("admin->users.go->FindAllUserInfo->Query :", Query)

	rows, err := db.Query(Query)
	if err != nil {
		log.Print("admin->users.go->FindAllUserInfo=->db.Query Failed:", err)
		return []Users{}, err
	}
	defer rows.Close()

	/*得出查询结果*/
	for rows.Next() {

		/*定义一个临时用户信息*/
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
			log.Print("admin->users.go->FindAllUserInfo->rows.Scan Failed:", err)
			continue
		}

		log.Print("admin->users.go->FindAllUserInfo->rows.Scan->tmpusr ", tmpusr)

		alluser = append(alluser, tmpusr)
	}
	return alluser, nil
}

func (users *Users) AddOneUser(db *sql.DB) (int, error) {

	Query := fmt.Sprintf(StmtAddOneUser, users.Username, users.Password, users.DefaultSchema)
	log.Print("admin->users.go->AddOneUser->Query :", Query)

	res, err := db.Exec(Query)
	if err != nil {
		log.Print("admin->users.go->AddOneUser->db.Exec Failed:", err)
		return 1, err //add user failed
	}

	rowsAffected, err := res.RowsAffected()
	log.Print("admin->users.go->AddOneUser->RowsAffected: ", rowsAffected)

	cmd.LoadUserToRuntime(db)
	cmd.SaveUserToDisk(db)
	return 0, nil
}

func (users *Users) DeleteOneUser(db *sql.DB) (int, error) {

	Query := fmt.Sprintf(StmtDeleteOneUser, users.Username)
	log.Print("admin->users.go->DeleteOneUser->Query:", Query)

	res, err := db.Exec(Query)
	if err != nil {
		log.Print("admin->users.go->DeleteOneUser->db.Exec Failed:", err)
		return 1, err //delte failed
	}

	rowsAffected, err := res.RowsAffected()
	log.Print("admin->users.go->DeleteOneUser->RowsAffected:", rowsAffected)

	cmd.LoadUserToRuntime(db)
	cmd.SaveUserToDisk(db)
	return 0, nil //delete success

}

// 更新一个用户所有信息，使用PATCH方法
func (users *Users) UpdateOneUserInfo(db *sql.DB) (int, error) {

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
		users.Username)

	log.Print("users->UpdateOneUserInfo->st: ", Query)

	res, err := db.Exec(Query)
	if err != nil {
		log.Print("admin->users.go->UpdateOneUserInfo->db.Exec Failed :", err)
		return 1, err
	}

	rowsAffected, err := res.RowsAffected()
	log.Print("admin->users.go->UpdateOneUser->RowsAffected: ", rowsAffected)

	cmd.LoadUserToRuntime(db)
	cmd.SaveUserToDisk(db)

	return 0, nil
}
