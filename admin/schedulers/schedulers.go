package schedulers

import (
	"database/sql"
	"fmt"
	"log"
	"proxysql-master/admin/cmd"
)

type Schedulers struct {
	Id         int64  `json:"id" db:"id"`
	Active     int64  `json:"active" db:"active"`
	IntervalMs int64  `json:"interval_ms" db:"interval_ms"`
	FileName   string `json:"filename" db:"filename"`
	Arg1       string `json:"arg1" db:"arg1"`
	Arg2       string `json:"arg2" db:"arg2"`
	Arg3       string `json:"arg3" db:"arg3"`
	Arg4       string `json:"arg4" db:"arg4"`
	Arg5       string `json:"arg5" db:"arg5"`
	Comment    string `json:"comment" db:"comment"`
}

const (
	/*添加一个新的调度器*/
	StmtAddOneScheduler = `
	INSERT 
	INTO 
		scheduler(filename,interval_ms) 
	VALUES(%q,%d)`

	/*删除一个调度器*/
	StmtDeleteOneScheduler = `
	DELETE 
	FROM 
		scheduler 
	WHERE id = %d`

	/*更新一个调度器*/
	StmtUpdateOneScheduler = `
	UPDATE 
		scheduler 
	SET 
		active = %d,
		interval_ms=%d,
		filename = %q,
		arg1=%q,
		arg2=%q,
		arg3=%q,
		arg4=%q,
		arg5=%q,
		comment=%q 
	WHERE 
		id = %d`

	/*查询所有调度器*/
	StmtFindAllScheduler = `
	SELECT 
		id,
		active,
		interval_ms,
		filename,
		ifnull(arg1,""),
		ifnull(arg2,""),
		ifnull(arg3,""),
		ifnull(arg4,""),
		ifnull(arg5,""),
		comment 
	FROM 
		scheduler 
	LIMIT %d 
	OFFSET %d`
)

//查找出所有定时器
func (schld *Schedulers) FindAllSchedulerInfo(db *sql.DB, limit int64, skip int64) ([]Schedulers, error) {

	/*定义保存调度器的变量*/
	var allscheduler []Schedulers

	Query := fmt.Sprintf(StmtFindAllScheduler, limit, skip)
	log.Print("admin->scheduler->FindAllSchedulerInfo->Query :", Query)

	rows, err := db.Query(Query)
	if err != nil {
		log.Print("admin->scheduler.go->FindAllSchedulerInfo->err: ", err)
		return []Schedulers{}, err
	}
	defer rows.Close()

	/*得出结果*/
	for rows.Next() {

		var tmpscheduler Schedulers

		err = rows.Scan(
			&tmpscheduler.Id,
			&tmpscheduler.Active,
			&tmpscheduler.IntervalMs,
			&tmpscheduler.FileName,
			&tmpscheduler.Arg1,
			&tmpscheduler.Arg2,
			&tmpscheduler.Arg3,
			&tmpscheduler.Arg4,
			&tmpscheduler.Arg5,
			&tmpscheduler.Comment,
		)

		if err != nil {
			log.Print("admin->scheduler.go->FindAllSchedulerInfo db.Query Failed: ", err)
			continue
		}

		log.Print("admin->scheduler.go->FindAllScheduler->tmpscheduler", tmpscheduler)
		allscheduler = append(allscheduler, tmpscheduler)
	}

	return allscheduler, nil
}

//添加一个新调度器
func (schld *Schedulers) AddOneScheduler(db *sql.DB) (int, error) {

	Query := fmt.Sprintf(StmtAddOneScheduler, schld.FileName, schld.IntervalMs)
	log.Print("admin-scheduler-AddOneScheduler->st: ", Query)

	res, err := db.Exec(Query)
	if err != nil {
		log.Print("admin-schedulers.go->AddOneScheduler->db.Exec Failed:", err)
		return 1, err
	}

	rowsAffected, err := res.RowsAffected()
	log.Print("admin-scheduler.go->AddOneScheduler->rowsAffected:", rowsAffected)

	cmd.LoadSchedulerToRuntime(db)
	cmd.SaveSchedulerToDisk(db)

	return 0, nil
}

//删除一个调度器
func (schld *Schedulers) DeleteOneScheduler(db *sql.DB) (int, error) {

	Query := fmt.Sprintf(StmtDeleteOneScheduler, schld.Id)
	log.Print("admin-scheduler->DeleteOneScheduler->st: ", Query)

	res, err := db.Exec(Query)
	if err != nil {
		log.Print("admin-scheduler->DeleteOneScheduler->db.Exec Failed: ", err)
		return 1, err
	}

	rowsAffected, err := res.RowsAffected()
	log.Print("admin-scheduler->DeleteOneScheduler->RowsAffected: ", rowsAffected)

	cmd.LoadSchedulerToRuntime(db)
	cmd.SaveSchedulerToDisk(db)
	return 0, nil
}

//更新一个调度器
func (schld *Schedulers) UpdateOneSchedulerInfo(db *sql.DB) (int, error) {

	Query := fmt.Sprintf(StmtUpdateOneScheduler, schld.Active, schld.IntervalMs, schld.FileName, schld.Arg1, schld.Arg2, schld.Arg3, schld.Arg4, schld.Arg5, schld.Comment, schld.Id)
	log.Print("admin-scheduler->UpdateOneScheduler->st: ", Query)

	res, err := db.Exec(Query)
	if err != nil {
		log.Print("admin-scheduler->UpdateOneSchedulerInfo->db.Exec err:", err)
		return 1, err
	}

	rowsAffected, err := res.RowsAffected()
	log.Print("admin-scheduler.go-UpdateOneScheduler->rowsAffected", rowsAffected)

	cmd.LoadSchedulerToRuntime(db)
	cmd.SaveSchedulerToDisk(db)

	return 0, nil
}
