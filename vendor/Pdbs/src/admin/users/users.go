package users

import (
	"database/sql"
	"log"
)

type (
	Users struct {
		Id           int64  `json:"id" db:"id"`
		AccessHost   string `json:"access_host" db:"access_host"`
		UserName     string `json:"username" db:"username"`
		Password     string `json:"password" db:"password"`
		AccessSchema string `json:"access_schema" db:"access_schema"`
		Active       int64  `json:"active" db:"active"`
		CreatedTime  string `json:"created_time" db:"created_time"`
		UpdatedTime  string `json:"updated_time" db:"updated_time"`
	}
)

const (
	StmtSelectAllUsers = `SELECT * FROM t_users`
	StmtCreateOneUser  = `INSERT INTO t_users(access_host,username,password,active) VALUES(?,?,?,?)`
	StmtUpdateOneUser  = `UPDATE t_users SET access_host=?,username=?,password=?,access_schema=?,active=? WHERE id=?`
	StmtDeleteOneUser  = `DELETE FROM t_users WHERE id=?`
)

/*查询所有用户信息*/
func (usr *Users) SelectAllUsers(db *sql.DB) ([]Users, error) {

	/*定义一个数组用于保存返回值*/
	var usrs []Users

	/*执行查询操作*/
	st, err := db.Query(StmtSelectAllUsers)
	if err != nil {
		log.Print("SelectAllUser db.Query Failed: ", err)
		return []Users{}, err
	}
	defer st.Close()
	log.Print("SelectAllUsers db.Query Success")

	for st.Next() {
		var tmpusr Users

		err := st.Scan(
			&tmpusr.Id,
			&tmpusr.AccessHost,
			&tmpusr.UserName,
			&tmpusr.Password,
			&tmpusr.AccessSchema,
			&tmpusr.Active,
			&tmpusr.CreatedTime,
			&tmpusr.UpdatedTime,
		)

		if err != nil {
			log.Print("SelectAllUser st.Scan Failed: ", err)
			continue
		}

		log.Print("SelectAllUsers tmpusr = ", tmpusr)
		usrs = append(usrs, tmpusr)
	}

	return usrs, nil
}

/*新建一个用户*/
func (usr *Users) CreateOneUser(db *sql.DB) (int, error) {
	st, err := db.Prepare(StmtCreateOneUser)
	if err != nil {
		log.Print("CreateOneUser db.Prepare Failed: ", err)
		return 1, err
	}

	log.Print("CreateOneUser db.Prepare Success")

	res, err := st.Exec(
		usr.AccessHost,
		usr.UserName,
		usr.Password,
		usr.Active,
	)
	if err != nil {
		log.Print("CreateOneUser st.Exec Failed:", err)
		return 1, err
	}

	log.Print("CreateOneUser st.Exec Success")

	/*获取影响行数*/
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Print("CreateOneUser res.RowsAffected Failed: ", err)
		return 2, err
	}

	log.Print("CreateOneUser RowsAffected: ", rowsAffected)
	return 0, nil
}

/*更新一个用户信息*/
func (usr *Users) UpdateOneUser(db *sql.DB) (int, error) {
	st, err := db.Prepare(StmtUpdateOneUser)
	if err != nil {
		log.Print("UpdateOneUser db.Prepare Failed", err)
		return 1, err
	}

	res, err := st.Exec(
		usr.AccessHost,
		usr.UserName,
		usr.Password,
		usr.AccessSchema,
		usr.Active,
		usr.Id,
	)
	if err != nil {
		log.Print("UpdateOneUser st.Exec Failed: ", err)
		return 1, err
	}

	/*获取影响行数*/
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Print("UpdateOneUser res.RowsAffected Failed: ", err)
		return 2, err
	}

	log.Print("UpdateOneUser res.RowsAffected :", rowsAffected)
	return 0, nil
}

/*删除一个用户*/
func (usr *Users) DeleteOneUser(db *sql.DB) (int, error) {
	st, err := db.Prepare(StmtDeleteOneUser)
	if err != nil {
		log.Print("DeleteOneUser db.Prepare Failed: ", err)
		return 1, err
	}

	log.Print("DeleteOneUser db.Prepare Success")

	res, err := st.Exec(usr.Id)
	if err != nil {
		log.Print("DeleteOneUser st.Exec Failed: ", err)
		return 1, err
	}

	/*获取影响行数*/
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Print("DeleteOneUser res.RowsAffected Failed: ", err)
		return 2, err
	}

	log.Print("DeleteOneUser res.RowsAffected: ", rowsAffected)
	return 0, nil
}
