package schedulers

import (
	"database/sql"
	"fmt"
	"log"
	"proxysql-master/admin/cmd"
)

type Schedulers struct {
	Id         int64  `json:"id" db:"id"`
	Active     int64  `json:"id" db:"id"`
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
	StmtAddOneScheduler    = `INSERT INTO scheduler(filename,interval_ms) values(%d,%q)`
	StmtDeleteOneScheduler = `DELETE FROM scheduler WHERE id = %d`
	StmtUpdateOneScheduler = `UPDATE scheduler SET active = %d,interval_ms=%d,filename = %q,arg1 = %q,arg2=%q,arg3=%q,arg4=%q,arg5=%q,comment=%q WHERE id = %d`
	StmtFindAllScheduler   = `SELECT id,active,interval_ms,filename,ifnull(arg1,""),ifnull(arg2,""),ifnull(arg3,""),ifnull(arg4,""),ifnull(arg5,""),comment FROM scheduler limit %d offset %d`
)

//查找出所有定时器
func FindAllSchedulerInfo(db *sql.DB, limit int64, skip int64) []Schedulers {
	var allscheduler []Schedulers
	var tmpscheduler Schedulers
	st := fmt.Sprintf(StmtFindAllScheduler, limit, skip)
	log.Print("scheduler->FindAllSchedulerInfo->st :", st)
	rows, err := db.Query(st)
	if err != nil {
		log.Print("scheduler.go->FindAllSchedulerInfo->err: ", err)
	}
	defer rows.Close()

	for rows.Next() {
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
		allscheduler = append(allscheduler, tmpscheduler)
	}
	return allscheduler
}

//添加一个新调度器
func (schld *Schedulers) AddOneScheduler(db *sql.DB) int {
	st := fmt.Sprintf(StmtAddOneScheduler, schld.FileName, schld.IntervalMs)
	log.Print("scheduler-AddOneScheduler->st: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("schedulers.go->AddOneScheduler->db.Query :", err)
		return 1
	}
	cmd.LoadSchedulerToRuntime(db)
	cmd.SaveSchedulerToDisk(db)
	log.Print("scheduler->AddOneScheduler->db.Query Success")
	return 0
}

//删除一个调度器
func (schld *Schedulers) DeleteOneScheduler(db *sql.DB) int {
	st := fmt.Sprintf(StmtDeleteOneScheduler, schld.Id)
	log.Print("scheduler->DeleteOneScheduler->st: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("scheduler->DeleteOneScheduler->db.Query: ", err)
		return 1
	}
	cmd.LoadSchedulerToRuntime(db)
	cmd.SaveSchedulerToDisk(db)
	return 0
}

//更新一个调度器
func (schld *Schedulers) UpdateOneSchedulerInfo(db *sql.DB) int {
	st := fmt.Sprintf(StmtUpdateOneScheduler, schld.Active, schld.IntervalMs, schld.FileName, schld.Arg1, schld.Arg2, schld.Arg3, schld.Arg4, schld.Arg5, schld.Comment, schld.Id)
	log.Print("scheduler->UpdateOneScheduler->st: ", st)
	_, err := db.Query(st)
	if err != nil {
		log.Print("scheduler->UpdateOneSchedulerInfo->db.Query err:", err)
		return 1
	}
	cmd.LoadSchedulerToRuntime(db)
	cmd.SaveSchedulerToDisk(db)
	return 0
}
